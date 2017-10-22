# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

"""Manages the generation and uploading of Python wheel CIPD packages."""

import collections
import glob
import itertools
import os
import shutil
import sys
import tempfile

from . import cipd
from . import source
from . import platform
from . import util


UniversalSpec = collections.namedtuple('UniversalSpec', (
    'pyversions'))


_Spec = collections.namedtuple('_Spec', (
    'name', 'version', 'universal'))
class Spec(_Spec):

  @property
  def tuple(self):
    return (self.name, self.version)

  @property
  def tag(self):
    return '%s-%s' % (self.name, self.version)


_Wheel = collections.namedtuple('_Wheel', (
    'spec', 'plat', 'pyversion', 'filename'))
class Wheel(_Wheel):

  @property
  def pyversion_str(self):
    if self.spec.universal:
      pyv = self.spec.universal.pyversions
      if pyv is None:
        return 'py2.py3'
      assert 'py2' in pyv
      return 'py2'
    return 'cp%s' % (self.pyversion,)

  @property
  def abi(self):
    if self.spec.universal or not self.plat.wheel_abi:
      return 'none'
    return self.plat.wheel_abi

  @property
  def platform(self):
    return ['any'] if self.spec.universal else self.plat.wheel_plat

  @property
  def primary_platform(self):
    """The platform to use when naming intermediate wheels and requesting
    wheel from "pip". Generally, platforms that this doesn't work on (e.g.,
    ARM) will not have wheels in PyPi, and platforms with wheels in
    PyPi will have only one platform.

    This is also used for naming when building wheels; this choice is
    inconsequential in this context, as the wheel is renamed after the build.
    """
    return self.platform[0]

  def default_filename(self):
    d = {
        'name': self.spec.name.replace('-', '_'),
        'version': self.spec.version,
        'pyversion': self.pyversion_str,
        'abi': self.abi,
        'platform': '.'.join(self.platform),
    }
    return '%(name)s-%(version)s-%(pyversion)s-%(abi)s-%(platform)s.whl' % d

  def path(self, system):
    return os.path.join(system.wheel_dir, self.filename)

  def cipd_package(self, templated=False):
    base_path = ['infra', 'python', 'wheels']
    if self.spec.universal:
      base_path += ['%s-%s' % (self.spec.name, self.pyversion_str)]
    else:
      base_path += [self.spec.name]
      if not templated:
        base_path += [
          '%s_%s_%s' % (self.plat.cipd_platform, self.pyversion_str, self.abi)]
      else:
        base_path += ['${vpython_platform}']

    tags = [
      'version:%s' % (self.spec.version,),
    ]
    return cipd.Package(
      name=('/'.join(p.replace('.', '_') for p in base_path)).lower(),
      tags=tuple(tags),
      install_mode=cipd.INSTALL_SYMLINK,
      compress_level=cipd.COMPRESS_NONE,
    )


class PlatformNotSupported(Exception):
  """Exception raised by Builder.build when the specified wheel's platform is
  not support."""


class Builder(object):

  def __init__(self, spec, build_fn, arch_map=None, abi_map=None,
               only_plat=None, skip_plat=None):
    """Initializes a new wheel Builder.

    spec (Spec): The wheel specification.
    build_fn (callable): Callable build function, used to generate the acutal
        wheel.
    arch_map (dict or None): Naming map for architectures. If the current
        platform has an entry in this map, the generated wheel will use the
        value as the "platform" field.
    abi_map (dict or None): Naming map for ABI. If the current platform
        has an entry in this map, the generated wheel will use the
        value as the "abi" field.
    only_plat (iterable or None): If not None, this Builder will only declare
        that it can build for the named platforms.
    skip_plat (iterable or None): If not None, this Builder will avoid declaring
        that it can build for the named platforms.
    """

    self._spec = spec
    self._build_fn = build_fn
    self._arch_map = arch_map or {}
    self._abi_map = abi_map or {}
    self._only_plat = frozenset(only_plat or ())
    self._skip_plat = frozenset(skip_plat or ())

  @property
  def spec(self):
    return self._spec

  def wheel(self, _system, plat):
    wheel = Wheel(
        spec=self._spec,
        plat=plat,
        # Only support Python 2.7 for now, can augment later.
        pyversion='27',
        filename=None)

    # Determine our package's wheel filename. This incorporates "abi" and "arch"
    # override maps, which are a priori knowledge of the package repository's
    # layout. This can differ from the local platform value if the package was
    # valid and built for multiple platforms, which seems to happen on Mac a
    # lot.
    plat_wheel = wheel._replace(
      plat=wheel.plat._replace(
        wheel_abi=self._abi_map.get(plat.name, plat.wheel_abi),
        wheel_plat=self._arch_map.get(plat.name, plat.wheel_plat),
      ),
    )
    return wheel._replace(
        filename=plat_wheel.default_filename(),
    )

  def supported(self, plat):
    if self._only_plat and plat.name not in self._only_plat:
      return False
    if plat.name in self._skip_plat:
      return False
    return True

  def build(self, wheel, system, rebuild=False):
    if not self.supported(wheel.plat):
      raise PlatformNotSupported()

    pkg_path = os.path.join(system.pkg_dir, '%s.pkg' % (wheel.filename,))
    if not rebuild and os.path.isfile(pkg_path):
      util.LOGGER.info('Package is already built: %s', pkg_path)
      return pkg_path

    # Rebuild the wheel, if necessary.
    wheel_path = wheel.path(system)
    if rebuild or not os.path.isfile(wheel_path):
      self._build_fn(system, wheel)
    else:
      util.LOGGER.info('Wheel is already built: %s', wheel_path)

    # Create a CIPD package for the wheel.
    util.LOGGER.info('Creating CIPD package: %r => %r', wheel_path, pkg_path)
    with system.temp_subdir('cipd_%s_%s' % wheel.spec.tuple) as tdir:
      shutil.copy(wheel_path, tdir)
      system.cipd.create_package(wheel.cipd_package(), tdir, pkg_path)

    return pkg_path


def check_run(system, dx, work_root, cmd, cwd=None):
  """Runs a command |cmd|.

  Args:
    system (runtime.System): The System instance.
    dx (dockcross.Image or None): The DockCross image to use. If None, the
        command will be run on the local system.
    work_root (str): The work root directory. If |dx| is not None, this will
        be the directory mounted as "/work" in the Docker environment.
    cmd (list): The command to run. Any components that are paths beginning
        with |work_root| will be automatically made relative to |work_root|.
    cwd (str or None): The working directory for the command. If None,
        |work_root| will be used. Otherwise, |cwd| must be a subdirectory of
        |work_root|.
    """
  if dx is None:
    if cmd[0] == 'python':
      cmd[0] = sys.executable
    return system.check_run(cmd, cwd=cwd or work_root)
  return dx.check_run(work_root, cmd, cwd=cwd)


def check_run_script(system, dx, work_root, script, args=None, cwd=None):
  """Runs a script, |script|.

  An anonymous file will be created under |work_root| holding the specified
  script.

  Args:
    script (list): A list of script lines to execute.
    See "check_run" for full argument definition.
  """
  with util.anonfile(work_root, text=True) as fd:
    for line in script:
      fd.write(line)
      fd.write('\n')
  os.chmod(fd.name, 0755)

  util.LOGGER.debug('Running script (path=%s): %s', fd.name, script)
  cmd = [fd.name]
  if args:
    cmd.extend(args)
  return check_run(system, dx, work_root, cmd, cwd=cwd)


def _stage_wheel_for_package(system, wheel_dir, wheel):
  # Find the wheel in "wheel_dir". We scan expecting exactly one wheel.
  wheels = glob.glob(os.path.join(wheel_dir, '*.whl'))
  assert len(wheels) == 1, 'Unexpected wheels: %s' % (wheels,)
  dst = os.path.join(system.wheel_dir, wheel.filename)

  source_path = wheels[0]
  util.LOGGER.debug('Identified source wheel: %s', source_path)
  shutil.copy(source_path, dst)


def _build_package(system, wheel):
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    check_run(
        system,
        None,
        tdir,
        [
          'python', '-m', 'pip', 'download',
          '--no-deps',
          '--only-binary=:all:',
          '--abi=%s' % (wheel.abi,),
          '--python-version=%s' % (wheel.pyversion,),
          '--platform=%s' % (wheel.primary_platform,),
          '%s==%s' % (wheel.spec.name, wheel.spec.version),
        ],
        cwd=tdir)

    _stage_wheel_for_package(system, tdir, wheel)


def _build_source(system, wheel, src, universal=False):
  dx = system.dockcross_image(wheel.plat)
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    build_dir = system.repo.ensure(src, tdir)

    bdist_wheel_opts = []
    if universal:
      bdist_wheel_opts.append('--universal')
    else:
      bdist_wheel_opts.append('--plat-name=%s' % (wheel.primary_platform,))

    cmd = [
      'python', '-m', 'pip', 'wheel',
      '--no-deps',
      '--only-binary=:all:',
      '--wheel-dir', tdir,
    ]
    for opt in bdist_wheel_opts:
      cmd += ['--build-option=%s' % (opt,)]
    cmd.append('.')

    check_run(
        system,
        dx,
        tdir,
        cmd,
        cwd=build_dir)

    _stage_wheel_for_package(system, tdir, wheel)


def _build_cryptography(system, wheel, src, openssl_src):
  dx = system.dockcross_image(wheel.plat)
  assert dx, 'Docker image required for compilation.'
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    # Unpack "cryptography".
    crypt_dir = system.repo.ensure(src, tdir)

    # Unpack "OpenSSL" into the "openssl/" subdirectory.
    openssl_dir = system.repo.ensure(openssl_src, tdir)

    # Build OpenSSL. We build this out of "openssl_dir" and install to
    # <openssl_dir>/PREFIX, so that will be the on-disk path to our OpenSSL
    # libraries.
    #
    # "Configure" must be run in the directory in which it builds, so we
    # `cd` into "openssl_dir" using dockcross "run_args".
    prefix = dx.workpath('prefix')
    check_run_script(
        system,
        dx,
        tdir,
        [
          '#!/bin/bash',
          'set -e',
          'export NUM_CPU="$(getconf _NPROCESSORS_ONLN)"',
          'echo "Using ${NUM_CPU} CPU(s)"',
          ' '.join([
            './Configure',
            '-fPIC',
            '--prefix=%s' % (prefix,),
            'no-shared',
            'no-ssl3',
            wheel.plat.openssl_target,
          ]),
          'make -j${NUM_CPU}',
          'make install',
        ],
        cwd=openssl_dir,
    )

    # Build "cryptography".
    d = {
      'prefix': prefix,
    }
    check_run_script(
        system,
        dx,
        tdir,
        [
          '#!/bin/bash',
          'set -e',
          'export CFLAGS="' + ' '.join([
            '-I%(prefix)s/include' % d,
            '$CFLAGS',
          ]) + '"',
          'export LDFLAGS="' + ' '.join([
            '-L%(prefix)s/lib' % d,
            '-L%(prefix)s/lib64' % d,
            '$LDFLAGS',
          ]) + '"',
          ' '.join([
            'python2.7',
            'setup.py',
            'build_ext',
            '--include-dirs', '/usr/cross/include',
            '--library-dirs', '/usr/cross/lib',
            '--force', 'build',
            '--force', 'build_scripts',
            '--executable=/usr/local/bin/python',
            '--force', 'bdist_wheel',
            '--plat-name', wheel.primary_platform,
          ]),
        ],
        cwd=crypt_dir,
    )

    _stage_wheel_for_package(system, os.path.join(crypt_dir, 'dist'), wheel)

def _build_opencv(system, wheel, numpy_version):
  # See "resources/build-opencv.sh" for more information.
  opencv_python = (
      'infra/third_party/source/opencv_python_repo',
      'git_revision:83b0ac8a200195d466bd7b4b5ac26923c98f0a64')
  virtualenv = (
        'infra/python/virtualenv',
        'version:15.1.0')

  # Build our "numpy" wheel.
  numpy_builder = BuildWheel('numpy', numpy_version)
  numpy_wheel = numpy_builder.wheel(system, wheel.plat)
  numpy_builder.build(numpy_wheel, system)
  numpy_path = numpy_wheel.path(system)

  dx = system.dockcross_image(wheel.plat)
  with system.temp_subdir('%s_%s' % wheel.spec.tuple) as tdir:
    # Copy external resources into "tdir" (workdir).
    script_path = util.copy_to(util.resource_path('build-opencv.sh'), tdir)
    numpy_path = util.copy_to(numpy_path, tdir)

    # Get OpenCV source and check out the correct version.
    opencv_path = os.path.join(tdir, 'opencv_cipd')
    system.cipd.install(opencv_python[0], opencv_python[1], opencv_path)
    opencv_path = util.copy_to(opencv_path, os.path.join(tdir, 'opencv'))

    # Get VirtualEnv source.
    venv_root = os.path.join(tdir, 'virtualenv')
    system.cipd.install(virtualenv[0], virtualenv[1], venv_root)
    venv_path = os.path.join(venv_root, 'virtualenv-15.1.0')

    # Run our build script.
    workdir = util.ensure_directory(tdir, 'workdir')
    check_run(
        system,
        dx,
        tdir,
        [
          'sh',
          script_path,
          workdir,
          opencv_path,
          wheel.spec.version,
          venv_path,
          numpy_path,
        ],
    )

    _stage_wheel_for_package(
        system, os.path.join(workdir, 'wheel', 'dist'), wheel)


def BuildWheel(name, version, **kwargs):
  """General-purpose wheel builder.

  If the wheel is "packaged" (see arg for description), it is expected that it
  is resident in PyPi and will be downloaded; otherwise, it will be built from
  source.

  Args:
    name (str): The wheel name.
    version (str): The wheel version.
    packaged (iterable or None): The names of platforms that have this wheel
        available via PyPi. If None, a default set of packaged wheels will be
        generated based on standard PyPi expectations, encoded with each
        Platform's "packaged" property.
    kwargs: Keyword arguments forwarded to Builder.

  Returns (Builder): A configured Builder for the specified wheel.
  """
  pypi_src = source.pypi_sdist(name, version)
  spec = Spec(name=name, version=pypi_src.version, universal=None)

  packaged = set(kwargs.pop('packaged', (p.name for p in platform.PACKAGED)))

  def build_fn(system, wheel):
    if wheel.plat.name in packaged:
      return _build_package(system, wheel)
    return _build_source(system, wheel, pypi_src)

  return Builder(spec, build_fn, **kwargs)


def BuildCryptographyWheel(name, crypt_src, openssl_src, packaged=None,
                           arch_map=None):
  """Specialized wheel builder for the "cryptography" package.

  Args:
    name (str): The wheel name.
    crypt_src (Source): The Source for the cryptography package. The wheel
        version will be extracted from this.
    openssl_src (Source): The OpenSSL source to build against.
    packaged (iterable or None): The names of platforms that have this wheel
        available via PyPi. If None, a default set of packaged wheels will be
        generated based on standard PyPi expectations, encoded with each
        Platform's "packaged" property.
    arch_map: (See Builder's "arch_map" argument.)

  Returns (Builder): A configured Builder for the specified wheel.
  """
  spec = Spec(name=name, version=crypt_src.version, universal=None)

  def build_fn(system, wheel):
    if wheel.plat.name in (packaged or ()):
      return _build_package(system, wheel)
    return _build_cryptography(system, wheel, crypt_src, openssl_src)

  return Builder(spec, build_fn, arch_map=arch_map)


def BuildOpenCVWheel(name, version, numpy_version, packaged=None,
                     arch_map=None, only_plat=None, skip_plat=None):
  """Specialized wheel builder for the "OpenCV" package.

  Args:
    name (str): The wheel name.
    version (str): The OpenCV version (must be a Git tag within the source).
    numpy_version (str): The "numpy" wheel version to build against.
        version will be extracted from this.
    packaged (iterable or None): The names of platforms that have this wheel
        available via PyPi. If None, will build from source for all platforms.
    arch_map: (See Builder's "arch_map" argument.)
    only_plat: (See Builder's "only_plat" argument.)
    skip_plat (iterable or None): If not None, this Builder will avoid declaring
        that it can build for the named platforms.

  Returns (Builder): A configured Builder for the specified wheel.
  """
  spec = Spec(name=name, version=version, universal=None)

  def build_fn(system, wheel):
    if wheel.plat.name in (packaged or ()):
      return _build_package(system, wheel)
    return _build_opencv(system, wheel, numpy_version)

  return Builder(spec, build_fn, arch_map=arch_map, only_plat=only_plat,
                 skip_plat=skip_plat)



def Packaged(name, version, only_plat, **kwargs):
  """Wheel builder for prepared wheels that must be downloaded from PyPi.

  Args:
    name (str): The wheel name.
    version (str): The wheel version.
    only_plat: (See Builder's "only_plat" argument.)
    kwargs: Keyword arguments forwarded to Builder.

  Returns (Builder): A configured Builder for the specified wheel.
  """
  spec = Spec(
      name=name,
      version=version,
      universal=None,
  )

  def build_fn(system, wheel):
    return _build_package(system, wheel)

  kwargs['only_plat'] = only_plat
  return Builder(spec, build_fn, **kwargs)


def Universal(name, version, pyversions=None, **kwargs):
  """Universal wheel version of BuildWheel.

  Args:
    name (str): The wheel name.
    version (str): The wheel version.
    pyversions (iterable or None): The list of "python" wheel fields (see
        "Wheel.pyversion_str"). If None, a default Python version will be used.
    kwargs: Keyword arguments forwarded to Builder.

  Returns (Builder): A configured Builder for the specified wheel.
  """
  spec = Spec(
      name=name,
      version=version,
      universal=UniversalSpec(
        pyversions=pyversions,
      ),
  )

  return Builder(spec, _build_package, **kwargs)


def UniversalSource(name, pypi_version, pyversions=None, pypi_name=None,
                    **kwargs):
  """Universal wheel version of BuildWheel that always builds from source.

  Args:
    name (str): The wheel name.
    version (str): The wheel version.
    pyversions (iterable or None): The list of "python" wheel fields (see
        "Wheel.pyversion_str"). If None, a default Python version will be used.
    pypi_name (str or None): Name of the package in PyPi. This can be useful
        when translating between the CIPD package name (uses underscores) and
        the PyPi package name (may use hyphens).
    kwargs: Keyword arguments forwarded to Builder.

  Returns (Builder): A configured Builder for the specified wheel.
  """
  pypi_src = source.pypi_sdist(
      name=pypi_name or name,
      version=pypi_version)

  spec = Spec(
      name=name,
      version=pypi_src.version,
      universal=UniversalSpec(
        pyversions=pyversions,
      ),
  )

  def build_fn(system, wheel):
    return _build_source(system, wheel, pypi_src, universal=True)

  return Builder(spec, build_fn, **kwargs)


SPECS = {s.spec.tag: s for s in (
  BuildWheel('coverage', '4.3.4'),
  BuildWheel('cffi', '1.10.0',
    arch_map={
      'mac-x64': 'macosx_10_6_intel',
    },
  ),
  BuildWheel('numpy', '1.11.3',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={
        'mac-x64': '.'.join([
          'macosx_10_6_intel',
          'macosx_10_9_intel',
          'macosx_10_9_x86_64',
          'macosx_10_10_intel',
          'macosx_10_10_x86_64',
        ]),
      },
  ),
  BuildWheel('numpy', '1.12.1',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={
        'mac-x64': '.'.join([
          'macosx_10_6_intel',
          'macosx_10_9_intel',
          'macosx_10_9_x86_64',
          'macosx_10_10_intel',
          'macosx_10_10_x86_64',
        ]),
      },
      skip_plat=('linux-arm64',),
  ),

  BuildWheel('psutil', '5.2.2',
      abi_map={
        'windows-x86': 'none',
        'windows-x64': 'none',
      },
      arch_map={
        'mac-x64': '.'.join([
          'macosx_10_6_intel',
          'macosx_10_9_intel',
          'macosx_10_9_x86_64',
          'macosx_10_10_intel',
          'macosx_10_10_x86_64',
        ]),
      },
      packaged=['windows-x86', 'windows-x64'],
  ),

  Packaged('scipy', '0.19.0',
      ['mac-x64', 'manylinux-x86', 'manylinux-x64'],
      arch_map={
        'mac-x64': '.'.join([
          'macosx_10_6_intel',
          'macosx_10_9_intel',
          'macosx_10_9_x86_64',
          'macosx_10_10_intel',
          'macosx_10_10_x86_64',
        ]),
      },
  ),

  BuildOpenCVWheel('opencv_python', '2.4.13.2', '1.11.3',
      skip_plat=['windows-x86', 'windows-x64', 'mac-x64']),

  BuildOpenCVWheel('opencv_python', '3.2.0.7', '1.12.1',
      packaged=[
        'mac-x64',
        'manylinux-x86',
        'manylinux-x64',
        'windows-x86',
        'windows-x64',
      ],
      arch_map={
        'mac-x64': '.'.join([
          'macosx_10_6_intel',
          'macosx_10_9_intel',
          'macosx_10_9_x86_64',
          'macosx_10_10_intel',
          'macosx_10_10_x86_64',
        ]),
      },
  ),

  BuildCryptographyWheel('cryptography',
      source.pypi_sdist('cryptography', '2.0.3'),
      source.remote_archive(
          name='openssl',
          version='1.1.0f',
          url='https://www.openssl.org/source/openssl-1.1.0f.tar.gz',
      ),
      arch_map={
        'mac-x64': 'macosx_10_6_intel',
      },
      packaged=[
        'manylinux-x86',
        'manylinux-x64',
        'mac-x64',
        'windows-x86',
        'windows-x64',
      ],
  ),

  BuildWheel('crcmod', '1.7', packaged=()),
  BuildWheel('grpcio', '1.4.0'),
  BuildWheel('scan-build', '2.0.8', packaged=()),

  # Prefer to use 'cryptography' instead of PyCrypto, if possible. We have to
  # use PyCrypto for GAE dev server (it's the only crypto package available on
  # GAE). Since we support it only on Linux and OSX, build only for these
  # platforms.
  BuildWheel('pycrypto', '2.6.1',
      packaged=(),
      only_plat=['manylinux-x64', 'mac-x64'],
  ),

  Universal('appdirs', '1.4.3'),
  UniversalSource('Appium_Python_Client', '0.24',
                   pypi_name='Appium-Python-Client'),
  Universal('asn1crypto', '0.22.0'),
  Universal('astunparse', '1.5.0'),
  Universal('boto', '2.48.0'),
  Universal('Django', '1.9'),
  Universal('enum34', '1.1.6', pyversions=['py2', 'py3']),
  Universal('funcsigs', '1.0.2'),
  UniversalSource('google_compute_engine', '2.6.2',
                  pypi_name='google-compute-engine'),
  Universal('google_api_python_client', '1.6.2'),
  UniversalSource('apache-beam', '2.0.0'),
  UniversalSource('httplib2', '0.10.3'),
  Universal('idna', '2.5'),
  Universal('ipaddress', '1.0.18', pyversions=['py2']),
  Universal('mock', '2.0.0'),
  Universal('oauth2client', '4.0.0'),
  Universal('packaging', '16.8'),
  Universal('pbr', '3.0.0'),
  Universal('protobuf', '3.2.0'),
  Universal('pyasn1', '0.2.3'),
  Universal('pyasn1_modules', '0.0.8'),
  UniversalSource('pycparser', '2.17'),
  Universal('pyopenssl', '17.2.0'),
  Universal('pyparsing', '2.2.0'),
  Universal('requests', '2.13.0'),
  Universal('rsa', '3.4.2'),
  Universal('selenium', '3.4.1'),
  Universal('setuptools', '34.3.2'),
  Universal('six', '1.10.0'),
  Universal('uritemplate', '3.0.0'),
)}
SPEC_NAMES = sorted(SPECS.keys())

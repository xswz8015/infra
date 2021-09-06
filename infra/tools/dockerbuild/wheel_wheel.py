# Copyright 2017 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from . import build_platform
from . import source
from . import util

from .builder import Builder, BuildPackageFromPyPiWheel, BuildPackageFromSource

from .build_types import Spec


class SourceOrPrebuilt(Builder):

  def __init__(self,
               name,
               version,
               pyversions=None,
               default=True,
               patches=(),
               patch_base=None,
               patch_version=None,
               **kwargs):
    """General-purpose wheel builder.

    If the wheel is "packaged" (see arg for description), it is expected that it
    is resident in PyPi and will be downloaded; otherwise, it will be built from
    source.

    Args:
      name (str): The wheel name.
      version (str): The wheel version.
      pyversions (iterable or None): The list of "python" wheel fields (see
          "Wheel.pyversion_str"). If None, a default Python version will be
          used.
      default (bool): If true, the wheel will be built by default.
      patches (tuple): Short patch names to apply to the source tree.
      patch_base (str or None): Optionally override the base names for patches.
      patch_version (str or None): If set, this string is appended to the CIPD
          version tag, for example if set to 'chromium.1', the version tag
          for version 1.2.3 of the wheel would be 'version:1.2.3.chromium.1'.
      packaged (iterable or None): The names of platforms that have this wheel
          available via PyPi. If None, a default set of packaged wheels will be
          generated based on standard PyPi expectations, encoded with each
          Platform's "packaged" property.
      env (Dict[str, str]|None): Envvars to set when building the wheel from
          source.
      kwargs: Keyword arguments forwarded to Builder.
    """
    self._pypi_src = source.pypi_sdist(
        name, version, patches=patches, patch_base=patch_base)
    self._packaged = set(
        kwargs.pop('packaged', (p.name for p in build_platform.PACKAGED)))
    self._env = kwargs.pop('env', None)
    version_suffix = '.' + patch_version if patch_version else None

    super(SourceOrPrebuilt, self).__init__(
        Spec(
            name,
            self._pypi_src.version,
            universal=False,
            pyversions=pyversions,
            default=default,
            version_suffix=version_suffix), **kwargs)

  def build_fn(self, system, wheel):
    if wheel.plat.name in self._packaged:
      return BuildPackageFromPyPiWheel(system, wheel)
    return BuildPackageFromSource(system, wheel, self._pypi_src, self._env)


class MultiWheel(Builder):

  def __init__(self,
               name,
               version,
               wheels,
               pyversions=None,
               only_plat=None,
               skip_plat=None,
               default=True):
    """Builds a wheel consisting of multiple other wheels.

    Bundles can be useful when a user always wants a common set of packages.

    Args:
      name (str): The name of the bundle wheel.
      version (str): The bundle wheel version. Note there is no patch_version
          for MultiWheels, as the bundle wheel version can be set to any value
          we like.
      wheels (iterable): A set of embedded wheel rules to add to the bundle.
      pyversions (iterable or None): The list of "python" wheel fields (see
          "Wheel.pyversion_str"). If None, a default Python version will be
          used.
      default (bool): If true, the wheel will be built by default.
      only_plat: (See Builder's "only_plat" argument.)
      skip_plat: (See Builder's "skip_plat" argument.)
    """
    self._wheels = wheels
    super(MultiWheel, self).__init__(
        Spec(
            name,
            version,
            universal=False,
            pyversions=pyversions,
            default=default,
            version_suffix=None),
        only_plat=only_plat,
        skip_plat=skip_plat)

  def build_fn(self, system, wheel):
    sub_wheels = []
    for w in self._wheels:
      subwheel_plat = (
          build_platform.ALL[build_platform.UNIVERSAL[0]]
          if w.spec.universal else wheel.plat)
      sub_wheel = w.wheel(system, subwheel_plat)
      util.LOGGER.info('Building sub-wheel: %s', sub_wheel)
      # Any time the MultiWheel is built, rebuild all the subwheels.
      sub_wheels += w.build_wheel(sub_wheel, system, rebuild=True)
    return sub_wheels


class Prebuilt(Builder):
  """Wheel builder for prepared wheels that must be downloaded from PyPi.

  Args:
    name (str): The wheel name.
    version (str): The wheel version.
    only_plat: (See Builder's "only_plat" argument.)
    pyversions (iterable or None): The list of "python" wheel fields (see
        "Wheel.pyversion_str"). If None, a default Python version will be
        used.
    default (bool): If true, the wheel will be built by default.
    kwargs: Keyword arguments forwarded to Builder.
  """

  def __init__(self,
               name,
               version,
               only_plat,
               pyversions=None,
               default=True,
               **kwargs):
    kwargs['only_plat'] = only_plat
    super(Prebuilt, self).__init__(
        Spec(
            name,
            version,
            universal=False,
            pyversions=pyversions,
            default=default,
            version_suffix=None), **kwargs)

  def build_fn(self, system, wheel):
    return BuildPackageFromPyPiWheel(system, wheel)


class Universal(Builder):

  def __init__(self, name, version, pyversions=None, default=True, **kwargs):
    """Universal wheel version of SourceOrPrebuilt.

    Args:
      name (str): The wheel name.
      version (str): The wheel version.
      pyversions (iterable or None): The list of "python" wheel fields (see
          "Wheel.pyversion_str"). If None, a default Python version will be
          used.
      kwargs: Keyword arguments forwarded to Builder.
    """
    super(Universal, self).__init__(
        Spec(
            name,
            version,
            universal=True,
            pyversions=pyversions,
            default=default,
            version_suffix=None,
        ), **kwargs)

  def build_fn(self, system, wheel):
    return BuildPackageFromPyPiWheel(system, wheel)


class UniversalSource(Builder):

  def __init__(self,
               name,
               pypi_version,
               pyversions=None,
               pypi_name=None,
               patches=(),
               patch_base=None,
               patch_version=None,
               **kwargs):
    """Universal wheel version of SourceOrPrebuilt that always builds from
    source.

    Args:
      name (str): The wheel name.
      version (str): The wheel version.
      pyversions (iterable or None): The list of "python" wheel fields (see
          "Wheel.pyversion_str"). If None, a default Python version will be
          used.
      pypi_name (str or None): Name of the package in PyPi. This can be useful
          when translating between the CIPD package name (uses underscores) and
          the PyPi package name (may use hyphens).
      patches (tuple): Short patch names to apply to the source tree.
      patch_base (str or None): Optionally override the base name for patches.
      patch_version (str or None): If set, this string is appended to the CIPD
          version tag, for example if set to 'chromium.1', the version tag
          for version 1.2.3 of the wheel would be 'version:1.2.3.chromium.1'.
          May not be used in combination with 'patches', since 'patches'
          already appends a hash of the patches to the version.
      kwargs: Keyword arguments forwarded to Builder.

    Returns (Builder): A configured Builder for the specified wheel.
    """
    self._pypi_src = source.pypi_sdist(
        name=pypi_name or name,
        version=pypi_version,
        patches=patches,
        patch_base=patch_base)
    if patches:
      if patch_version:
        raise ValueError('patches and patch_version may not be used together.')
      # For backward compatibility.
      version_suffix = '-' + self._pypi_src.patches_hash
    elif patch_version:
      version_suffix = '.' + patch_version
    else:
      version_suffix = None
    super(UniversalSource, self).__init__(
        Spec(
            name,
            self._pypi_src.version,
            universal=True,
            pyversions=pyversions,
            default=True,
            version_suffix=version_suffix,
        ), **kwargs)

  def build_fn(self, system, wheel):
    return BuildPackageFromSource(system, wheel, self._pypi_src)

  def md_data_fn(self):
    if not self._pypi_src.patches:
      return []

    return ['\n* custom patches: %s' % (', '.join(self._pypi_src.patches),)]

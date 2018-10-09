# Copyright 2018 The Chromium Authors. All rights reserved.
# Use of this source code is governed by a BSD-style license that can be
# found in the LICENSE file.

from recipe_engine.recipe_api import Property

DEPS = [
  "recipe_engine/buildbucket",
  "recipe_engine/cipd",
  "recipe_engine/file",
  "recipe_engine/path",
  "recipe_engine/platform",
  "recipe_engine/properties",
  "recipe_engine/raw_io",
  "recipe_engine/runtime",
  "recipe_engine/step",

  "third_party_packages_ng",
]

PROPERTIES = {
  'GOOS': Property(),
  'GOARCH': Property(),
  'package_prefix': Property(kind=str, default=''),
  'load_dupe': Property(kind=bool, default=False),
}

def RunSteps(api, GOOS, GOARCH, package_prefix, load_dupe):
  builder = api.path['cache'].join('builder')
  api.third_party_packages_ng.package_prefix = package_prefix

  api.step('echo package_prefix', [
    'echo', api.third_party_packages_ng.package_prefix])

  # do a checkout in `builder`
  pkgs = api.third_party_packages_ng.load_packages_from_path(
    builder.join('package_repo'))

  # For the test, also explicitly build 'tool@1.5.0-rc1', which should de-dup
  # with the default tool@latest.
  pkgs.add('tool@1.5.0-rc1')

  # doing it twice should raise a DuplicatePackage exception
  if load_dupe:
    api.third_party_packages_ng.load_packages_from_path(
      builder.join('dup_repo'))

  _, unsupported = api.third_party_packages_ng.ensure_uploaded(
    pkgs, '%s-%s' % (GOOS, GOARCH))

  excluded = {'unsupported'}
  if api.platform.is_win:
    excluded.add('posix_tool')
  assert unsupported == excluded, "unexpected: %r" % (unsupported,)

  # doing it again should hit caches
  api.third_party_packages_ng.ensure_uploaded(pkgs, '%s-%s' % (GOOS, GOARCH))


def GenTests(api):
  pkgs = sorted(dict(
    bottom_dep='''
    create {
      source { cipd { pkg: "source/bottom_dep" default_version: "1.0" } }
      build {}
    }
    upload { pkg_prefix: "prefix/deps" }
    ''',
    tool='''
    create {
      source {
        git {
          repo: "https://go.repo/tool"
          tag_pattern: "v%s"
          version_join: "."
        }
        subdir: "src/go.repo/tool"
        patch_dir: "patches"
        patch_version: "chops.1"
      }
      build {
        # We use an older version of the tool to bootstrap new versions.
        tool: "tool@0.9.0"
        dep: "bottom_dep"

        install: "install.sh"
        install: "intel"
      }
      package {
        version_file: ".versions/tool.cipd_version"
      }
    }

    create {
      platform_re: "mac-.*"
      build {
        install: "install-mac.sh"
      }
      package {
        install_mode: symlink
      }
      verify {
        test: "test.py"
        test: "mac"
      }
    }

    create {
      platform_re: "windows-.*"
      verify {
        test: "test.py"
        test: "windows"
      }
    }

    create {
      platform_re: "linux-.*"
      verify {
        test: "test.py"
        test: "linux"
      }
    }

    create {
      platform_re: "linux-arm.*"
      build {
        install: "install.sh"
        install: "arm"
      }
    }

    create {
      platform_re: "linux-amd64"
      build {
        # on linux-amd64 we self-bootstrap the tool
        tool: ""  # clears tool@0.9.0
        install: "install_bootstrap.sh"
      }
    }

    upload { pkg_prefix: "prefix/build_tools" }
    ''',

    deep_dep='''
    create {
      source { cipd {pkg: "source/deep_dep" default_version: "1.0.0"} }
    }
    upload { pkg_prefix: "prefix/deps" }
    ''',

    dep='''
    create {
      source { cipd {pkg: "source/dep" default_version: "1.0.0"} }
      build {
        tool: "tool"
        dep: "deep_dep"
      }
    }
    upload { pkg_prefix: "prefix/deps" }
    ''',

    pkg='''
    create {
      source { script { name: "fetch.py" } }
      build {
        tool: "tool"
        dep:  "dep"
      }
    }
    upload { pkg_prefix: "prefix/tools" }
    ''',

    unsupported='''
    create { unsupported: true }
    ''',

    posix_tool='''
    create {
      platform_re: "linux-.*|mac-.*"
      source {
        cipd { pkg: "source/posix_tool" default_version: "1.2.0" }
        unpack_archive: true
      }
      build {}  # default build options
    }
    upload { pkg_prefix: "tools" }
    ''',

    already_uploaded='''
    create {
      source { cipd { pkg: "source/already_uploaded" default_version: "1.5.0-rc1" } }
    }
    upload { pkg_prefix: "tools" }
    ''',

    # This doesn't have a 'build' step. It just fetches something (e.g. gcloud
    # SDK), and then re-uploads it.
    fetch_and_package='''
    create { source { script { name: "fetch.py" } } }
    upload {
      pkg_prefix: "tools"
      universal: true
    }
    ''',
  ).items())

  def mk_name(*parts):
    return '.'.join(parts)

  for goos, goarch in (('linux', 'amd64'),
                       ('linux', 'armv6l'),
                       ('windows', 'amd64'),
                       ('mac', 'amd64')):
    plat_name = 'win' if goos == 'windows' else goos

    sep = '\\' if goos == 'windows' else '/'
    pkg_repo_path = sep.join(
      ['[CACHE]', 'builder', 'package_repo', '%s', '3pp.pb'])
    plat = '%s-%s' % (goos, goarch)

    test = (api.test('integration_test_%s-%s' % (goos, goarch))
      + api.runtime(is_luci=True, is_experimental=True)
      + api.platform(plat_name, 64)  # assume all hosts are 64 bits.
      + api.properties(GOOS=goos, GOARCH=goarch)
      + api.buildbucket.ci_build()
      + api.step_data('find package specs',
                      api.file.glob_paths([
                        pkg_repo_path % name for name, _ in pkgs]))
      + api.override_step_data(mk_name(
        'building already_uploaded',
        'cipd describe experimental/3pp/tools/already_uploaded/%s' % plat
      ), api.cipd.example_describe(
        'experimental/3pp/tools/already_uploaded/%s' % plat,
        version='version:1.5.0-rc1', test_data_tags=['version:1.5.0-rc1']))
    )

    if plat_name != 'win':
      # posix_tool says it needs an archive unpacking.
      test += api.step_data(mk_name(
        'building posix_tool', 'fetch sources', 'unpack_archive',
        'find archive to unpack',
      ), api.file.glob_paths(['archive.tgz']))

    for pkg, spec in pkgs:
      test += api.step_data("load package specs.read '%s'" % pkg,
                            api.file.read_text(spec))
    yield test

  yield (api.test('empty spec')
      + api.properties(GOOS='linux', GOARCH='amd64')
      + api.step_data(
          'find package specs',
          api.file.glob_paths(['[CACHE]/builder/package_repo/bad/3pp.pb']))
  )

  yield (api.test('bad spec')
      + api.properties(GOOS='linux', GOARCH='amd64')
      + api.step_data(
          'find package specs',
          api.file.glob_paths(['[CACHE]/builder/package_repo/bad/3pp.pb']))
      + api.step_data(
          "load package specs.read 'bad'", api.file.read_text('narwhal'))
      + api.expect_exception('BadParse')
  )

  yield (api.test('duplicate load')
      + api.properties(GOOS='linux', GOARCH='amd64', load_dupe=True)
      + api.step_data(
          'find package specs',
          api.file.glob_paths(['[CACHE]/path/something/3pp.pb']))
      + api.step_data(
          "load package specs.read 'something'",
          api.file.read_text('create {}'))
      + api.step_data(
          'find package specs (2)',
          api.file.glob_paths(['[CACHE]/path/something/3pp.pb']))
      + api.expect_exception('DuplicatePackage')
  )

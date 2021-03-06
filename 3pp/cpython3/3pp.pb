create {
  verify { test: "python_test.py" }
  source { patch_version: "chromium.22" }
  package { version_file: ".versions/cpython3.cipd_version" }
}

create {
  platform_re: "linux-.*|mac-.*"
  source {
    git {
      repo: "https://chromium.googlesource.com/external/github.com/python/cpython"
      tag_pattern: "v%s",

      # Pin to 3.8.x for now.
      version_restriction: { op: LT val: "3.10a0"}
    }
    patch_dir: "patches"
  }
  build {
    # no binutils on mac since it includes some tools like 'ar' that we don't
    # actually want
    tool: "build_support/pip_bootstrap"
    tool: "tools/autoconf"
    tool: "tools/sed"
  }
}

create {
  platform_re: "mac-.*"
  build {
    dep: "static_libs/bzip2"
    dep: "static_libs/libuuid"
    dep: "static_libs/ncursesw"
    dep: "static_libs/openssl"
    dep: "static_libs/readline"
    dep: "static_libs/sqlite"
    dep: "static_libs/xzutils"
    dep: "static_libs/zlib"
  }
}

create {
  platform_re: "linux-.*"
  build {
    dep: "static_libs/bzip2"
    dep: "static_libs/libffi"
    dep: "static_libs/libuuid"
    dep: "static_libs/ncursesw"
    dep: "static_libs/openssl"
    dep: "static_libs/readline"
    dep: "static_libs/sqlite"
    dep: "static_libs/xzutils"
    dep: "static_libs/zlib"

    # On Linux, we need to explicitly build libnsl; on other platforms, it is
    # part of 'libc'.
    dep: "static_libs/nsl"

    tool: "build_support/pip_bootstrap"
    tool: "tools/autoconf"
    tool: "tools/binutils"
    tool: "tools/sed"
  }
}

create {
  platform_re: "linux-arm.*|linux-mips.*"
  build {
    tool: "build_support/pip_bootstrap"
    tool: "tools/autoconf"
    tool: "tools/binutils"
    tool: "tools/sed"            # Used by python's makefiles

    tool: "tools/cpython3"
  }
}

create {
  platform_re: "windows-.*"
  source { script { name: "fetch.py" } }
  build {
    tool: "build_support/pip_bootstrap"
    tool: "tools/lessmsi"

    install: "install_win.sh"
  }
  verify { test: "python_test.py" }
}

upload { pkg_prefix: "tools" }

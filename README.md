# lddc

![Build status](https://github.com/opslabhqx/lddc/actions/workflows/go-release.yml/badge.svg)
![Licence: MIT](https://img.shields.io/github/license/opslabhqx/lddc)

A simple script that collects and copies shared library dependencies of a specified binary to a given directory.

[![asciicast](https://asciinema.org/a/669670.svg)](https://asciinema.org/a/669670)

## Getting Started

### Prerequisites

- `ldd` binary installed (usually available by default on most Linux distributions)

### Installation

To install `lddc`, you can use Homebrew. First, tap the repository (For Linux only):

```bash
brew tap opslabhqx/lddc
```

Then, install `lddc` with:

```bash
brew install lddc
```

### Usage

To collect and copy the shared library dependencies of a binary, use the following command:

```bash
./lddc <path to the binary> <path to copy the dependencies>
```

`<path to the binary>`: The path to the binary file whose dependencies you want to collect.
`<path to copy the dependencies>`: The directory where you want to copy the shared library dependencies.

### Example

```bash
% lddc $(which curl) lib

No such directory lib, creating...
Collecting the shared library dependencies for /usr/bin/curl...
Copying the dependencies to lib
Copying /lib/x86_64-linux-gnu/libcurl.so.4 to lib
Copying /lib/x86_64-linux-gnu/libz.so.1 to lib
Copying /lib/x86_64-linux-gnu/libc.so.6 to lib
Copying /lib/x86_64-linux-gnu/libnghttp2.so.14 to lib
Copying /lib/x86_64-linux-gnu/libidn2.so.0 to lib
Copying /lib/x86_64-linux-gnu/librtmp.so.1 to lib
Copying /lib/x86_64-linux-gnu/libssh.so.4 to lib
Copying /lib/x86_64-linux-gnu/libpsl.so.5 to lib
Copying /lib/x86_64-linux-gnu/libssl.so.3 to lib
Copying /lib/x86_64-linux-gnu/libcrypto.so.3 to lib
Copying /lib/x86_64-linux-gnu/libgssapi_krb5.so.2 to lib
Copying /lib/x86_64-linux-gnu/libldap.so.2 to lib
Copying /lib/x86_64-linux-gnu/liblber.so.2 to lib
Copying /lib/x86_64-linux-gnu/libzstd.so.1 to lib
Copying /lib/x86_64-linux-gnu/libbrotlidec.so.1 to lib
Copying /lib/x86_64-linux-gnu/libunistring.so.5 to lib
Copying /lib/x86_64-linux-gnu/libgnutls.so.30 to lib
Copying /lib/x86_64-linux-gnu/libhogweed.so.6 to lib
Copying /lib/x86_64-linux-gnu/libnettle.so.8 to lib
Copying /lib/x86_64-linux-gnu/libgmp.so.10 to lib
Copying /lib/x86_64-linux-gnu/libkrb5.so.3 to lib
Copying /lib/x86_64-linux-gnu/libk5crypto.so.3 to lib
Copying /lib/x86_64-linux-gnu/libcom_err.so.2 to lib
Copying /lib/x86_64-linux-gnu/libkrb5support.so.0 to lib
Copying /lib/x86_64-linux-gnu/libsasl2.so.2 to lib
Copying /lib/x86_64-linux-gnu/libbrotlicommon.so.1 to lib
Copying /lib/x86_64-linux-gnu/libp11-kit.so.0 to lib
Copying /lib/x86_64-linux-gnu/libtasn1.so.6 to lib
Copying /lib/x86_64-linux-gnu/libkeyutils.so.1 to lib
Copying /lib/x86_64-linux-gnu/libresolv.so.2 to lib
Copying /lib/x86_64-linux-gnu/libffi.so.8 to lib
Done!
```

# License

This project is licensed under the [MIT License](/LICENSE). See the LICENSE file for details.

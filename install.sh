#!/bin/bash

ARCH=$(uname -m)
OS=$(uname -s)

if [[ "$OS" != "Darwin" && "$OS" != "Linux" ]]; then
  echo "Error: unexpected OS ($OS), only Darwin and Linux are supported."
  exit 1;
fi;
if [[ "$ARCH" != "x86_64" && "$ARCH" != "arm64" ]]; then
  echo "Error: unexpected CPU arch ($ARCH), only x86_64 and arm64 are supported."
  exit 1;
fi;

URL="https://github.com/slasyz/mk/releases/download/latest/mk-$OS-$ARCH"
echo "-> Downloading $URL"
curl --location -o /tmp/mk-binary "$URL"
chmod +x /tmp/mk-binary
mv -n /usr/local/bin/mk /usr/local/bin/mk.bak 2>/dev/null || true
mv /tmp/mk-binary /usr/local/bin/mk

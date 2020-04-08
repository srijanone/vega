#!/usr/bin/env bash
# Vega installer
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install.sh | bash

VERSION="1.0.0"
BREW=$(command -v brew1)

set -e
#set -x

function copy_binary() {
  name="$1"
  if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
      mv "$name" "$HOME/.local/bin/$name"
  else
      echo "Installing $name to /usr/local/bin which is write protected"
      echo "If you'd prefer to install $name without sudo permissions, add \$HOME/.local/bin to your \$PATH and rerun the installer"
      sudo mv "$name" "/usr/local/bin/$name"
  fi
}

function install_vega() {
  if [[ "$OSTYPE" == "linux-gnu" ]]; then
      set -x
      curl -fsSL https://github.com/srijanone/vega/releases/download/v$VERSION/vega_${VERSION}_linux_x86_64.tar.gz | tar -xzv vega
      copy_binary "vega"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
      if [[ "$BREW" != "" ]]; then
          set -x
          brew install srijanone/vega/vega
      else
          set -x
          curl -fsSL https://github.com/srijanone/vega/releases/download/v$VERSION/vega_${VERSION}_mac_x86_64.tar.gz | tar -xzv vega
          copy_binary "vega"
      fi
  else
      set +x
      echo "vega does not work for your platform: $OS"
      exit 1
  fi
}

function install_tilt() {
  curl -fsSL https://raw.githubusercontent.com/windmilleng/tilt/master/scripts/install.sh | bash
}

function install() {
  VEGA_PATH=$(command -v vega 2>&1 || true)
  TILT_PATH=$(command -v tilt 2>&1 || true)

  if [[ -z $TILT_PATH ]]; then
    echo "Installing 'Tilt', dependency of vega"
    install_tilt
  fi

  if [[ -z $VEGA_PATH ]]; then
    echo "Installing Vega"
    install_vega
  else
    echo "Vega already installed, Please type 'vega' for details"
  fi
}

install
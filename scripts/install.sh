#!/usr/bin/env bash
# Vega Installer Script
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install.sh | bash

VERSION="1.0.6"
BREW=$(command -v brew)
BREW="" # Don't wanna use brew, as some people don't have brew or sudo permission

# Vega has dependency of tilt: tilt.dev
TILT_VERSION="0.14.3"

set -e

red="\033[31m"
green="\033[32m"
yellow="\033[33m"
blue="\033[34m"
bold="\033[1m"
reset="\033[0m"

show_next_step_message="NO"

function shell() {
  # ps -p $$ | tail -1 | awk '{print $NF}'
  echo "${SHELL##*/}"
}

function shell_rc() {
  local shell_name=$(shell)
  echo ".${shell_name/-/}rc"
}

function next_step_message() {
  local name="vega"
  echo -e "${yellow}${name} is installed to \$HOME/.local/bin, Please add following line to your $(shell_rc) file and reload it using: ${bold}source $(shell_rc) ${reset}"
  echo ""
  echo -e "${bold}export PATH=\"\$PATH:\$HOME/.local/bin\"${reset}"
  echo ""
}

function copy_binary() {
  name="$1"
  if [[ ":$PATH:" == *":$HOME/.local/bin:"* ]]; then
    mv "${name}" "$HOME/.local/bin/${name}"
  else
    show_next_step_message="YES"
    mkdir -p "$HOME/.local/bin"
    mv "${name}" "$HOME/.local/bin/${name}"
  fi
}

function install_vega() {
  if [[ "$OSTYPE" == "linux-gnu" ]]; then
    curl -fsSL https://github.com/srijanone/vega/releases/download/v${VERSION}/vega_linux_amd64.tar.gz | tar -xzv vega 2>/dev/null
    copy_binary "vega"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    if [[ "$BREW" != "" ]]; then
      brew install srijanone/vega/vega
    else
      curl -fsSL https://github.com/srijanone/vega/releases/download/v${VERSION}/vega_darwin_amd64.tar.gz | tar -xzv vega 2>/dev/null
      copy_binary "vega"
    fi
  else
    echo -e "${red}The Vega installer is not supported for your platform ${OS} ${reset}"
    echo -e "${red}Please file an issue at https://github.com/srijanone/vega/issues/new ${reset}"
    exit 1
  fi
}

function install_tilt() {
  if [[ "$OSTYPE" == "linux-gnu" ]]; then
    curl -fsSL https://github.com/tilt-dev/tilt/releases/download/v${TILT_VERSION}/tilt.${TILT_VERSION}.linux.x86_64.tar.gz | tar -xzv tilt 2>/dev/null
    copy_binary "tilt"
  elif [[ "$OSTYPE" == "darwin"* ]]; then
    if [[ "$BREW" != "" ]]; then
      brew install tilt-dev/tap/tilt
    else
      curl -fsSL https://github.com/tilt-dev/tilt/releases/download/v${TILT_VERSION}/tilt.${TILT_VERSION}.mac.x86_64.tar.gz | tar -xzv tilt 2>/dev/null
      copy_binary "tilt"
    fi
  else
    echo -e "${red}The Tilt installer is not supported for your platform: ${OS} ${reset}"
    echo -e "${red}Please file an issue at https://github.com/tilt-dev/tilt/issues/new ${reset}"
    exit 1
  fi
}

function install_git_secrets() {
  if [[ "$OSTYPE" == "linux-gnu" ]] || [[ "$OSTYPE" == "darwin"* ]]; then
    curl -sSL -o git-secrets -D - -L -s 'https://raw.githubusercontent.com/awslabs/git-secrets/master/git-secrets'
    chmod +x git-secrets
    copy_binary "git-secrets"
  else
    echo -e "${red}git-secrets installer is not supported for your platform: ${OS} ${reset}"
    echo -e "${red}Please file an issue at https://github.com/awslabs/git-secrets/issues/new ${reset}"
    exit 1
  fi
}

function install() {
  VEGA_PATH=$(command -v vega 2>&1 || true)
  TILT_PATH=$(command -v tilt 2>&1 || true)
  GIT_SECRETS_PATH=$(command -v git-secrets 2>&1 || true)

  if [[ -z $VEGA_PATH ]]; then
    echo -e "${green}Installing Vega${reset}"
    install_vega
  else  
    echo -e "${green}Vega already installed, Please run 'vega' for details${reset}"
  fi

  if [[ -z $TILT_PATH ]]; then
    echo -e "${green}Installing Dependencies${reset}"
    echo -e "${green}Installing Tilt${reset}"
    install_tilt
  fi

  if [[ -z $GIT_SECRETS_PATH ]]; then
    echo -e "${green}Installing git-secrets${reset}"
    install_git_secrets
  else
    echo -e "${green}git-secrets already installed, Please run 'git-secrets for details${reset}"
  fi

  if [[ "${show_next_step_message}" == "YES" ]]; then
    next_step_message
  fi
}

install

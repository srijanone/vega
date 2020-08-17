#!/usr/bin/env bash
# Git Secrets Installer Script
#
# Usage:
#   curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install_git_secrets.sh | bash

VERSION="1.0.4"

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
  local name="git-secrets"
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
  GIT_SECRETS_PATH=$(command -v git-secrets 2>&1 || true)

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

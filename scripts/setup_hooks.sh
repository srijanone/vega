#!/usr/bin/env bash
# Setup Git Hooks
#
# Usage:
#   ./setup_hooks.sh

VERSION="1.0.0"

GLOBAL_HOOKS_DIR="${HOME}/.git/hooks"

function setup_hooks() {
  printf "Setting up Git Hooks\n"

  printf "Creating Global Hooks Directory\n"
  mkdir -p "${GLOBAL_HOOKS_DIR}"

  printf "Setting Global Git Hooks: %s\n" "${GLOBAL_HOOKS_DIR}"
  git config --global core.hooksPath "${GLOBAL_HOOKS_DIR}"

  printf "Installing pre-commit hooks\n"
  cp ../hooks/generic/pre-commit/check-aws-credentials.sh "${GLOBAL_HOOKS_DIR}/pre-commit"
}

setup_hooks

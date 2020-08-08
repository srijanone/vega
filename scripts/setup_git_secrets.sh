#!/usr/bin/env bash
# Setup Git Secrets
#
# Usage:
#   ./setup_git_secrets.sh
#

function setup_git_secrets() {
  printf "Setting up Git Secrets\n"

  printf "Adding common AWS patterns to the git config...\n"
  git secrets --register-aws --global

  printf "Adding hooks to all local repositories...\n"
  git secrets --install -f ~/.git-templates/git-secrets
  git config --global init.templateDir ~/.git-templates/git-secrets

  printf "Registering Drupal secrets patters...\n"
  git secrets --add --global "(\"|')?(host|port|password|username)(\"|')?\s*(:|=>|=)\s*(\"|')?(".*")(\"|')?\s*"
}

setup_git_secrets

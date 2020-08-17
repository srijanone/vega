# Vega (वेग)

Vega (वेग) is developer friendly project scaffolding tool to speed up development process.

[![Release](https://github.com/srijanone/vega/workflows/Release/badge.svg)](https://github.com/srijanone/vega/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/srijanone/vega)](https://goreportcard.com/report/github.com/srijanone/vega)

## Installation

Several options to install:

- Via **Installer Script**: `curl -fsSL https://raw.githubusercontent.com/srijanone/vega/develop/scripts/install.sh | bash`
- Via **Homebrew**: `brew install srijanone/vega/vega`
  - Update vega: `brew update && brew upgrade vega`
  - For Mac and Linux
- Via **Go**: `go install github.com/srijanone/vega`
  - (This might install latest unreleases/bleeding-edge version)
- Via released binaries:
  - [releases](https://github.com/srijanone/vega/releases)

## Requirements

- git
- [tilt](https://docs.tilt.dev/install.html)
- [Docker](https://docs.docker.com/install/)
- [docker-compose](https://docs.docker.com/compose/install/)

---

## Getting Started

- `vega`: vega usage

- `vega init`: Initializes vega

- `vega starterkit list`: List all available starterkits

- `vega create my-drupal-app --starterkit drupal8-php-fpm-apache`: Bootload a new app using starterkit

- `vega up`: Get your docker containers up & running.

- `vega down`: Stop all docker containers.

The above commands are mostly used commands, please refer commands table for further details.

## Commands

| Command                                      | Description                                                                           | Arguments                                       | Output     |
| -------------------------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------- | ---------- |
| `vega`                                         | Prints out usage and help                                                             | \--home <path/to/home>                          |            |
| `vega version`                                 | Prints out version                                                                    |                                                 | Vega 1.0.0 |
| `vega home`                                    | Prints out home vega home                                                             |                                                 |            |
| `vega init`                                    | Initializes vega                                                                      |                                                 |            |
| `vega starterkit list`                         | List all available starterkits                                                        |                                                 | drupal9-php-fpm-apache<br>react    |
| `vega create [path] --starterkit <name>`       | Creates the starter kit at provided directory                                         | \--starterkit <name><br>\--repo <repo>          |            |
| `vega install [path]`                          | Install a starterkit to existing project                                              | \--repo <repo>                                   |            |
| `vega repo add <repo-name> <url>`              | Add another starterkit repo, Can choose local folder as well                          |                                                 |            |
| `vega repo list`                               | Lists all the repo available                                                          | \--repo <repo>                                  |            |
| `vega hooks install [path]`                    | Installs git hooks to specified path                                                  |                                                 |            |
| `vega up`                                      | Runs the application                                                                  | \--port <log-port><br>\--watch<br>\--no-browser |            |
| `vega down`                                    | Stops the application and deletes the resources                                       |                                                 |            |

#### Notes:
- All commands can take additional `--home` flag which will override default $VEGA_HOME
- `--repo` flag can take git url or local folder url
  - Examples:
    1. `vega repo add globe git@github.com:vs4vijay/vega-starterkits.git`
    2. `vega repo add new /Users/viz/SrijanX/custom`

---

## Development

- Run Vega: `go run main.go`
- Release Binaries: `goreleaser`
  - Github Token to be created and exported: `export GITHUB_TOKEN=<token>`
  - Make command is added to Makefile:
    - `make release-dry-run`        # to test and verify on local machine
    - `make release-using-gorelease`

---

## Release
- Releases are generated using Github Action Pipelines which runs `goreleaser`
- Create a new tag: `git tag origin v1.0.x`
- Push tag: `git push origin v1.0.x`

---

## Secrets
vega has been integrated with [git-secrets](https://github.com/awslabs/git-secrets) which adds following hooks to your repositories when ```vega hooks install``` is executed.

  1. ```pre-commit```: Used to check if any of the files changed in the commit
       use prohibited patterns.
  2. ```commit-msg```: Used to determine if a commit message contains a
       prohibited patterns.
  3. ```prepare-commit-msg```: Used to determine if a merge commit will
       introduce a history that contains a prohibited pattern at any point.
       Please note that this hook is only invoked for non fast-forward merges.

```vega hooks install``` overrides any current git hooks if you have added any. In case you would like to have multiple
hooks please refer: https://gist.github.com/carlos-jenkins/89da9dcf9e0d528ac978311938aade43

---

## Credits

- Inspiration from Draft (https://draft.sh)
- Tilt (https://tilt.dev) is used for running the applications
- git-secrets

## LICENSE

This software is covered under the Apache v2.0 license. You can read the [license here](LICENSE).

This software uses tilt binary, which is covered by the Apache v2.0 license too.
# वेग

Vega (वेग) is developer friendly project scaffolding tool to speed up development process.

![Release](https://github.com/srijanone/vega/workflows/Release/badge.svg)

## Installation

Several options to install:

- Via **Homebrew**: `brew install srijanone/vega/vega`
  - (Mac and Linux only)
  - Update vega: `brew update && brew upgrade vega` 
- Via **Go**: `go install github.com/srijanone/vega`
  - (This might install latest unreleases/bleeding-edge version)
- Via released binaries:
  - [releases](https://github.com/srijanone/vega/releases)

---

## Requirements

- git
- (tilt)[https://docs.tilt.dev/install.html]
- (Docker)[https://docs.docker.com/install/]

---

## Getting Started

- `vega init`: Initializes vega
![vega init](_screenshots/vega_init.png)
- `vega starterkit list`: List all available starterkits
![vega starterkit list](_screenshots/vega_starterkit_list.png)
- `vega create awesome-app --starterkit nodejs+redis`
![vega create](_screenshots/vega_create.png)
- `vega up`
- `vega down`
![vega down](_screenshots/vega_down.png)

---

## Commands

| Command                                      | Description                                                                           | Arguments                                       | Output     |
| -------------------------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------- | ---------- |
| `vega`                                         | Prints out usage and help                                                             | \--home <path/to/home>                          |            |
| `vega version`                                 | Prints out version                                                                    |                                                 | Vega 1.0.0 |
| `vega home`                                    | Prints out home vega home                                                             |                                                 |            |
| `vega init`                                    | Initializes vega                                                                      |                                                 |            |
| `vega starterkit list`                         | List all available starterkits                                                        |                                                 | drupal8<br>nodejs    |
| `vega create [path] --starterkit <name>`       | Creates the starter kit at provided directory                                         | \--starterkit <name><br>\--repo <repo>          |            |
| `vega repo add <repo-name> <url>`              | Add another starterkit repo, Can choose local folder as well                          |                                                 |            |
| `vega repo list`                               | Lists all the repo available                                                          |                                                 |            |
| `vega up`                                      | Runs the application                                                                  | \--port <log-port><br>\--watch<br>\--no-browser |            |
| `vega down`                                    | Remove the application resources                                                      |                                                 |            |

---

## Development
`go run main.go`

- For Releasing Binaries: `goreleaser`
  - Github Token to be created and exported: `export GITHUB_TOKEN=<token>`

---

## Credits

- Srijan Team (https://srijan.net)
- Inspiration from Draft (https://draft.sh)
- Utilized Tilt (https://tilt.dev) for running the application 


---

## LICENSE

This software is covered under the Apache v2.0 license. You can read the [license here](LICENSE).

This software uses tilt binary, which is covered by the Apache v2.0 license too.
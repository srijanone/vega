# वेग

Vega (वेग) is developer friendly project scaffolding tool to speed up development process.

![Release](https://github.com/srijanone/vega/workflows/Release/badge.svg)

## Installation

Via **Go Get**:

```
go get github.com/srijanone/vega
```

Via **Homebrew** (Mac only):

```
brew install srijanone/vega/vega
```

- Update vega: `brew update && brew upgrade vega` 

Via prebuilt binaries:
- [releases](https://github.com/srijanone/vega/releases)

---

## Getting Started

```console
vega init
vega starterkit list
vega create --starterkit <starterkit>
vega up
vega down
```

## Commands


| Command                                      | Description                                                                           | Arguments                                       | Output     |
| -------------------------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------- | ---------- |
| `vega`                                         | Prints out usage and help                                                             | `\--home <path/to/home>`                          |            |
| `vega version`                                 | Prints out version                                                                    |                                                 | Vega 1.0.0 |
| `vega home`                                    | Prints out home vega home                                                             |                                                 |            |
| `vega starterkit list`                         | List all available starterkits                                                        |                                                 | drupal8    |
| `vega create \[path\] --starterkit <name>`     | Creates the starter kit in current directory. Can choose folder/github repo. as well. | `\--starterkit <name><br>\--repo <repo>`          |            |
| `vega repo add <repo-name> <url>`              | Add another starterkit repo                                                           |                                                 |            |
| `vega repo list`                               | Lists all the repo available                                                          |                                                 |            |
| `vega up`                                      | Runs the application                                                                  | \--port <log-port><br>\--watch<br>\--no-browser |            |
| `vega down`                                    | Remove the application resources                                                      |                                                 |            |

---

### Development

`go run main.go`

- For Releasing Binaries: `goreleaser`

---

## To Do

- [x] Basic Functionality
- [x] Release Binaries
- [x] Dockerize
- [x] Makefile
  - run
  - test
  - clean
- [ ] Unit Tests

---

```



```
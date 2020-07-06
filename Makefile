PROJECT     := "vega"
PREFIX      := "/usr/local"

GIT_COMMIT  := `git rev-parse HEAD`
GIT_SHA     := `git rev-parse --short HEAD`
GIT_TAG     := `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "v0.0.0"`
BUILD_TIME  := `date +%s`


LDFLAGS     := ""
LDFLAGS     += -X=github.com/srijanone/vega/pkg/version.SemVer=$(GIT_TAG)
LDFLAGS     += -X=github.com/srijanone/vega/pkg/version.GitCommit=$(GIT_COMMIT)
LDFLAGS     += -X=github.com/srijanone/vega/pkg/version.BuildTime=$(BUILD_TIME)

OS          := `uname | tr '[:upper:]' '[:lower:]'`
OS_LIST     := darwin linux windows

ARCH        := `uname -m`
ARCH_LIST   := 386 amd64


.PHONY: info
info:
  @echo "info..."
  @echo "Version:           ${GIT_TAG}"
  @echo "Git Commit:        ${GIT_COMMIT}"
  @echo "Git SHA:           ${GIT_SHA}"
  @echo "Build Time:        ${BUILD_TIME}"
  @echo ""


.PHONY: fmt
fmt:
  gofmt -l -w .


.PHONY: run
run: build
  ./bin/$(PROJECT)


.PHONY: build
build: info
  CGO_ENABLED=0 go build -v -o bin/$(PROJECT) -ldflags "$(LDFLAGS)"


.PHONY: build-all
build-all:
  @for os in $(OS_LIST); do \
    for arch in $(ARCH_LIST); do \
      echo "Building for OS($${os}) and Arch($${arch})"; \
      GOOS=$${os} GOARCH=$${arch} go build -v -ldflags "$(LDFLAGS)" -o "bin/$(PROJECT)_$${os}_$${arch}"; \
    done \
  done


.PHONY: release-dry-run
release-dry-run:
  @echo "release dry run"
  goreleaser --snapshot --skip-publish --rm-dist


.PHONY: release-using-gorelease
release-using-gorelease:
  @echo "releasing..."
  goreleaser --rm-dist


.PHONY: clean
clean:
  @echo "cleaning..."
  GO111MODULE=on go clean --modcache
  rm -rf bin/


.PHONY: clean-home
clean-home:
  @echo "cleaning $VEGA_HOME..."
  rm -rf ~/.vega


.PHONY: install
install: build
  @echo "installing..."
  cp bin/$(PROJECT) $(PREFIX)/bin/$(PROJECT)


.PHONY: uninstall
uninstall:
  @echo "uninstalling..."
  rm -f $(PREFIX)/bin/$(PROJECT)

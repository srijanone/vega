PROJECT := "vega"


GIT_COMMIT 	:= `git rev-parse HEAD`
GIT_SHA 	:= `git rev-parse --short HEAD`
GIT_TAG 	:= `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary"`
BUILD_TIME  := `date -u +"%Y-%m-%dT%H:%M:%SZ"`

LDFLAGS := ""
LDFLAGS += -X=github.com/srijanone/vega/pkg/version.SemVer=$(GIT_TAG)
LDFLAGS += -X=github.com/srijanone/vega/pkg/version.GitCommit=$(GIT_COMMIT)
LDFLAGS += -X=github.com/srijanone/vega/pkg/version.BuildTime=$(BUILD_TIME)

OS 			:= `uname | tr '[:upper:]' '[:lower:]'`
ARCH 		:= `uname -m`


.PHONY: info
info:
	@echo "info..."


.PHONY: build
build: info
	GOOS=darwin GOARCH=386 go build -v -ldflags "$(LDFLAGS)"


.PHONY: clean
clean:
	@echo "cleaning..."
	rm -rf ~/.vega
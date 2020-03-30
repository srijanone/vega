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
OS_LIST		:= darwin linux windows

ARCH 		:= `uname -m`
ARCH_LIST	:= 386 amd64

.PHONY: info
info:
	@echo "info..."
	@echo "Version:       		${GIT_TAG}"
	@echo "Git Commit:       	${GIT_COMMIT}"
	@echo "Git SHA:       		${GIT_SHA}"
	@echo "Build Time:       	${BUILD_TIME}"
	@echo ""


.PHONY: fmt
fmt:
	gofmt -l -w .


.PHONY: build
build: info
	GOOS=darwin GOARCH=386 go build -v -ldflags "$(LDFLAGS)"


.PHONY: clean
clean:
	@echo "cleaning..."
	rm -rf ~/.vega
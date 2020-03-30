PROJECT := "vega"


GIT_COMMIT 	:= `git rev-parse HEAD`
GIT_SHA 	:= `git rev-parse --short HEAD`
GIT_TAG 	:= `git describe --tags --abbrev=0 --exact-match 2>/dev/null || echo "canary"`
BUILD_TIME  := `date -u +"%Y-%m-%dT%H:%M:%SZ"`



.PHONY: info
info:
	@echo "info..."


.PHONY: build
build: info
	go build -v


.PHONY: clean
clean:
	@echo "cleaning..."
	rm -rf ~/.vega
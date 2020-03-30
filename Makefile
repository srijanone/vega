PROJECT := "vega"





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
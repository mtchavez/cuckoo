ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover

default: test

deps:
	@go get -v $(GOPKGS)

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet test

test:
	@echo "[Test] running tests"
	go test -v -cover -bench=".*" -coverprofile=c.out

.PHONY: default test vet deps

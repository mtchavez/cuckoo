ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint \
		github.com/tools/godep

default: test

deps:
	@go get -u -v $(GOPKGS)
	@go get -u -v -t ./...
	@if [ `which godep` ] && [ -f ./Godeps/Godeps.json ]; then godep restore; fi

lint:
	@echo "[Lint] running golint"
	@golint

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet lint test

test:
	@echo "[Test] running tests"
	@if [ "$CI" ]; then goveralls -service=travis-ci; else go test -v -cover; fi

.PHONY: default golint test vet deps

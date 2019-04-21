ROOT := $(CURDIR)
GOPKGS = \
		golang.org/x/tools/cmd/cover \
		github.com/golang/lint/golint

default: test

deps:
	@go get -v $(GOPKGS)

lint:
	@echo "[Lint] running golint"
	@golint

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet lint test

test:
	@echo "[Test] running tests"
	go test -v -cover -bench=".*" -coverprofile=c.out

.PHONY: default golint test vet deps

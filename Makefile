ROOT := $(CURDIR)
GOPKGS = \
		github.com/axw/gocov/gocov \
		github.com/golang/lint/golint \
		github.com/golang/dep/cmd/dep

default: test

deps:
	@go get -u -v $(GOPKGS)
	@dep ensure

lint:
	@echo "[Lint] running golint"
	@golint

vet:
	@echo "[Vet] running go vet"
	@go vet

ci: deps vet lint test

test:
	@echo "[Test] running tests"
	gocov test > coverge/gocov.out; cat coverge/gocov.out | gocov report

.PHONY: default golint test vet deps

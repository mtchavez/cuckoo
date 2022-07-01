# Cuckoo Filter

[![Latest Version](http://img.shields.io/github/release/mtchavez/cuckoo.svg?style=flat-square)](https://github.com/mtchavez/cuckoo/releases)
[![Test](https://github.com/mtchavez/cuckoo/actions/workflows/test.yml/badge.svg)](https://github.com/mtchavez/cuckoo/actions/workflows/test.yml)
[![Go Documentation](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/mtchavez/cuckoo)
[![Go Report Card](https://goreportcard.com/badge/github.com/mtchavez/cuckoo)](https://goreportcard.com/report/github.com/mtchavez/cuckoo)
[![Maintainability](https://api.codeclimate.com/v1/badges/3e295a8cb3cfe6f8c1ee/maintainability)](https://codeclimate.com/github/mtchavez/cuckoo/maintainability)
[![Test Coverage](https://codecov.io/gh/mtchavez/cuckoo/branch/master/graph/badge.svg?token=5xaMOOsXEd)](https://codecov.io/gh/mtchavez/cuckoo)

Cuckoo Filter in Go

## Install

Install via `go get`

`go get -v github.com/mtchavez/cuckoo`

## Usage

- [New Filter](#new-filter)
- [Configuring](#configuration)
- [Insert](#insert)
- [Insert Unique](#insert-unique)
- [Lookup](#lookup)
- [Delete](#delete)
- [Item Count](#item-count)
- [Save](#save)
- [Load](#load)

### New Filter

Create a new filter with default configuration

```go
package main

import "github.com/mtchavez/cuckoo"

func main() {
  cuckoo.New()
}
```

### Configuration

You can configure a filter via a `ConfigOption` type and the composed config option
functions provided.

```go
package main

import "github.com/mtchavez/cuckoo"

func main() {
  options := []cuckoo.ConfigOption{
    cuckoo.BucketEntries(uint(24)),
    cuckoo.BucketTotal(uint(1 << 16)),
    cuckoo.FingerprintLength(uint(1)),
    cuckoo.Kicks(uint(250)),
  }
  cuckoo.New(options...)
}
```

### Insert

Inserting items into a filter

```go
package main

import "github.com/mtchavez/cuckoo"

func main() {
  filter := cuckoo.New()
  filter.Insert([]byte("special-items"))
}
```

### Insert Unique

Inserting items into a filter only if they do not already exist

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := cuckoo.New()
  filter.InsertUnique([]byte("special-items"))
  filter.InsertUnique([]byte("special-items"))
  if filter.ItemCount() != 1 {
    fmt.Println("Expected only 1 item")
  }
}
```

### Lookup

Check if items exist in the filter using Lookup

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := cuckoo.New()
  filter.Insert([]byte("special-items"))
  found := filter.Lookup([]byte("special-items"))
  if !found {
    fmt.Println("Expected to find item in filter")
  }
}
```

### Delete

Deleting an item if it exists in the filter

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := cuckoo.New()
  filter.Insert([]byte("special-items"))
  deleted := filter.Delete([]byte("special-items"))
  if !deleted {
    fmt.Println("Expected to delete item from filter")
  }
}
```

### Item Count

Getting the item count of filter. **Using Insert with duplicates will cause the
item count to be more like a total items inserted count**. Using InsertUnique
and checking the ItemCount will be more of a *real* item count.

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := cuckoo.New()
  filter.InsertUnique([]byte("special-items"))
  filter.InsertUnique([]byte("special-items"))
  if filter.ItemCount() != 1 {
    fmt.Println("Expected only 1 item")
  }
}
```

### Save

Encode and save a filter to be able to re-build or re-load back into memory.

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := New()
  item := []byte("Largo")
  filter.InsertUnique(item)
  filter.Save("./tmp/example_save.gob")
}
```

### Load

Load a filter back into memory from an encoded filter saved to a file.

```go
package main

import (
  "fmt"

  "github.com/mtchavez/cuckoo"
)

func main() {
  filter := New()
  item := []byte("Largo")
  filter.InsertUnique(item)
  filter.Save("./tmp/example_save.gob")

  loadedFilter, _ := Load("./tmp/example_save.gob")
  fmt.Printf("Loaded filter has same item? %v\n\n", loadedFilter.Lookup(item))
}
```

## Benchmarks

There are benchmark tests to check performance of the filter. The following results
were ran on a 2.3 GHz Intel Core i7

```
# Updated: 2022-07-01

goos: darwin
goarch: arm64
pkg: github.com/mtchavez/cuckoo

BenchmarkCuckooNew-10                 48            23354917 ns/op
BenchmarkInsert-10               3342568               806.5 ns/op
BenchmarkInsertUnique-10         6203035               194.7 ns/op
BenchmarkLookup-10               6465182               196.3 ns/op
```

## Tests

Run tests via `go test` or with provided `Makefile`

`go test -v -cover ./...` or `make test`

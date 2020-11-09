# bpool

[![ci](https://github.com/johejo/bpool/workflows/ci/badge.svg?branch=main)](https://github.com/johejo/bpool/actions?query=workflow%3Aci)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/johejo/bpool)](https://pkg.go.dev/github.com/johejo/bpool)
[![codecov](https://codecov.io/gh/johejo/bpool/branch/main/graph/badge.svg)](https://codecov.io/gh/johejo/bpool)
[![Go Report Card](https://goreportcard.com/badge/github.com/johejo/bpool)](https://goreportcard.com/report/github.com/johejo/bpool)

Package bpool is a minimal wrapper using sync.Pool for *bytes.Buffer.

## Example

```go
package bpool_test

import (
	"testing"

	"github.com/johejo/bpool"
)

var pool = bpool.New()

func Example() {
	buf := pool.Get()
	defer pool.Put(buf)
	for i := 0; i < 1024; i++ {
		buf.WriteString("hello")
	}
	println(buf.String())
}
```

## License

MIT

## Author

Mitsuo Heijo(@johejo)

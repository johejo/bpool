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

func TestPool(t *testing.T) {
	buf := pool.Get()
	defer pool.Put(buf)
	buf.WriteString("hello")
	if buf.String() != "hello" {
		t.Error("should be hello")
	}
}

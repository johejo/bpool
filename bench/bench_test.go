package bpool_test

import (
	"bytes"
	"testing"

	"github.com/johejo/bpool"
	"github.com/lestrrat-go/bufferpool"
	obpool "github.com/oxtoacart/bpool"
	"github.com/valyala/bytebufferpool"
)

const s = "1234567890-=qwertyuiop[]asdfghjkl;''zxcvbnm,./1234567890-=qwertyuiop[]asdfghjkl;''zxcvbnm,./1234567890-=qwertyuiop[]asdfghjkl;''zxcvbnm,./1234567890-=qwertyuiop[]asdfghjkl;''zxcvbnm,./"
const N = 1024 * 64

func Benchmark_no_pool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		for j := 0; j < N; j++ {
			buf.WriteString(s)
		}
		buf.Bytes()
	}
}

func Benchmark_pool0(b *testing.B) {
	pool := bpool.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := pool.Get()
		for j := 0; j < N; j++ {
			buf.WriteString(s)
		}
		buf.Bytes()
		pool.Put(buf)
	}
}

func Benchmark_bufferpool(b *testing.B) {
	pool := bufferpool.New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := pool.Get()
		for j := 0; j < N; j++ {
			buf.WriteString(s)
		}
		buf.Bytes()
		pool.Release(buf)
	}
}

func Benchmark_obool(b *testing.B) {
	pool := obpool.NewBufferPool(48)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := pool.Get()
		for j := 0; j < N; j++ {
			buf.WriteString(s)
		}
		buf.Bytes()
		pool.Put(buf)
	}
}

func Benchmark_bytebufferpool(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf := bytebufferpool.Get()
		for j := 0; j < N; j++ {
			buf.WriteString(s)
		}
		buf.Bytes()
		bytebufferpool.Put(buf)
	}
}

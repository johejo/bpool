// Package bpool is a minimal wrapper using sync.Pool for *bytes.Buffer.
package bpool

import (
	"bytes"
	"sync"
)

// Pool wraps a sync.Pool for *bytes.Buffer.
type Pool struct {
	pool sync.Pool
}

// Get returns an empty *bytes.Buffer.
func (p *Pool) Get() *bytes.Buffer {
	return p.pool.Get().(*bytes.Buffer)
}

// Put returns *bytes.Buffer to the pool.
func (p *Pool) Put(b *bytes.Buffer) {
	b.Reset()
	p.pool.Put(b)
}

// New returns a new *Pool.
func New() *Pool {
	return &Pool{
		pool: sync.Pool{
			New: func() interface{} {
				return new(bytes.Buffer)
			},
		},
	}
}

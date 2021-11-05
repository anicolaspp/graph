package graph

import (
	"sync"
)

type Pool struct {
	edges sync.Map
}

func NewPool() *Pool {
	return &Pool{
		edges: sync.Map{},
	}
}

// Get tries finding a matching edge for the given values and return it.
// A new edge is created and added to the pool when needed.
func (p *Pool) Get(a, b int) *E {
	e := &E{
		A: a,
		B: b,
	}

	str := e.String()
	v, _ := p.edges.LoadOrStore(str, e)
	return v.(*E)
}

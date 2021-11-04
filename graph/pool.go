package graph

import (
	"sync"
)

type Pool struct {
	mu sync.Mutex
	edges map[string]*E
}

func NewPool() *Pool {
	return &Pool{
		edges: map[string]*E{},
	}
}

// Get tries finding a matching edge for the given values and return it.
// A new edge is created and added to the pool when needed.
func (p *Pool) Get(a, b int) *E {
	e := &E{
		A: a,
		B: b,
	}
	
	p.mu.Lock()
	defer p.mu.Unlock()	
	
	str := e.String()
	v, ok := p.edges[str]
	if !ok {
		p.edges[str] = e
		return e
	}

	return v
}

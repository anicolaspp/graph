package graph

type Pool struct {
	edges map[string]*E
}

func NewPool() *Pool {
	return &Pool{
		edges: map[string]*E{},
	}
}

// Get try finding a matching edge for the given values and return it.
// A new edge is created and added to the pool when needed.
func (p *Pool) Get(a, b int) *E {
	e := &E{
		A: a,
		B: b,
	}

	str := e.String()
	v, ok := p.edges[str]
	if !ok {
		p.edges[str] = e
		return e
	}

	return v
}

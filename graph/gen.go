package graph

import (
	"fmt"
	"sort"
)

func Gen(size int) []*G {
	graphs := map[int][]*G{}
	graphs[1] = []*G{
		&G{
			Es: []E{{1, 1}},
		},
	}
	for i := 2; i <= size; i++ {
		for _, g := range graphs[i-1] {
			toAdd := add(g, i)
			graphs[i] = append(graphs[i], toAdd...)
		}
	}

	res := map[string]*G{}
	for _, g := range graphs[size] {
		m := map[string]E{}
		for _, e := range g.Es {
			if _, ok := m[e.String()]; !ok && e.String() != "(1,1)" {
				m[e.String()] = e
			}
		}
		gg := &G{}
		for _, e := range m {
			gg.Es = append(gg.Es, e)
		}
		res[gg.String()] = gg
	}

	xs := []*G{}
	for _, v := range res {
		xs = append(xs, v)
	}

	return xs
}

func add(g *G, n int) []*G {
	res := []*G{}
	for i := 1; i <= len(g.Nodes()); i++ {
		res = append(res, addE(g, n, i)...)
	}
	return res
}

func addE(g *G, n int, c int) []*G {
	f := &R{F: [][]int{}}
	Comb(g.Nodes(), c, []int{}, f)

	res := []*G{}
	for _, v := range f.F {
		edges := []E{}
		edges = append(edges, g.Es...)
		for _, node := range v {
			edges = append(edges, E{node, n})
		}

		res = append(res, &G{edges})
	}

	return res
}

func Comb(nodes List, n int, r List, f *R) {
	if len(r) == n {
		f.Add(r)
	} else {
		for i := 0; i < len(nodes); i++ {
			node := nodes[i]
			nodes = nodes.Remove(i)

			r = append(r, node)
			Comb(nodes, n, r, f)
			r = r.Remove(len(r) - 1)
			nodes = nodes.Insert(node, i)
		}
	}
}

type M map[string][]int

func (m M) String() string {
	s := "{"
	for k, v := range m {
		s += fmt.Sprintf("%v -> %v,", k, v)
	}
	s += "}"
	return s
}

type R struct {
	F [][]int
}

func (r *R) Add(x List) {
	r.F = append(r.F, x.Cp())
}

type List []int

func (l *List) String() string {
	cp := l.Cp()
	sort.Slice(cp, func(i int, j int) bool {
		return cp[i] <= cp[j]
	})
	s := ""
	for _, v := range cp {
		s = fmt.Sprintf("%v", v)
	}
	return s
}

func (l List) Cp() List {
	result := List{}
	for i := 0; i < len(l); i++ {
		result = append(result, l[i])
	}
	return result
}

// Remove removes item ith from l.
func (l List) Remove(i int) List {
	result := l[:i]
	if i+1 < len(l) {
		result = append(result, l[i+1:]...)
	}
	return result
}

// Insert inserts value v in index i on l.
func (l List) Insert(v int, i int) List {
	result := append(l[:i+1], l[i:]...)
	result[i] = v
	return result
}

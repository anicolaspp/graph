package graph

import (
	"fmt"
	"sort"
)

var (
	pool = NewPool()
)

// Gen generates all connected graphs of given size.
// G(n) = AllGraphsOfSize(n - 1) + add(n).
func Gen(size int) []*G {
	graphs := map[int][]*G{}
	graphs[1] = []*G{
		&G{
			Es: []*E{pool.Get(1, 1)},
		},
	}
	for i := 2; i <= size; i++ {
		ch := make(chan []*G)

		// For each graph of size n - 1, we can add a node N and generate
		// multiple new graphs of size n.
		//
		// If we have k graph of size n - 1, then we start k async processes
		// where each of them creates multiple graphs by adding new node N.
		n := i                // new node N we are adding.
		k := len(graphs[n-1]) // number of n - 1 graphs we have where we need to ad new node N.
		for j := 0; j < k; j++ {
			// k async processes.
			go genUsing(graphs[n-1][j], n, ch)

			// Setting the used value (graph) to nil improves memory.
			graphs[n-1][j] = nil
		}

		// Since the number of graphs exponentially increases, let's remove
		// the generation n - 1 since we don't need it any more, and save some
		// (a lot) of memory.
		delete(graphs, n-1)

		// Aggregates all graph from all parallel processes.
		// Reads output from the k async processes.
		for j := 0; j < k; j++ {
			toAdd := <-ch
			graphs[n] = append(graphs[n], toAdd...)
		}
	}

	ch := make(chan *G)
	for _, g := range graphs[size] {
		go clean(g, ch)
	}

	xs := []*G{}
	for range graphs[size] {
		cleaned := <-ch
		xs = append(xs, cleaned)
	}

	return xs
}

func genUsing(g *G, i int, ch chan<- []*G) {
	toAdd := add(g, i)
	ch <- toAdd
}

func clean(g *G, ch chan<- *G) {
	m := map[string]*E{}
	for _, e := range g.Es {
		if _, ok := m[e.String()]; !ok && e.String() != "(1,1)" {
			m[e.String()] = e
		}
	}
	gg := &G{}
	for _, e := range m {
		gg.Es = append(gg.Es, e)
	}
	ch <- gg
}

// add adds n to the graph g by using addE (adding all required edges).
func add(g *G, n int) []*G {
	res := []*G{}
	for i := 1; i <= len(g.Nodes()); i++ {
		res = append(res, addE(g, n, i)...)
	}
	return res
}

// addE adds all the new edges from the new node n to the existing graph g.
func addE(g *G, n int, c int) []*G {
	f := &R{F: [][]int{}}
	Comb(g.Nodes(), c, []int{}, f)

	res := []*G{}
	for _, v := range f.F {
		edges := []*E{}
		edges = append(edges, g.Es...)
		for _, node := range v {
			edges = append(edges, pool.Get(node, n))
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

// R.F contains all possible ways to select (n, x) where n is the
// nodes (1..n), and x is increasing from (1..n).
// Comb([1,2], 1) -> [[1],[2]]
// Comb([1,2], 2) -> [[1,2]]
// Comb([1,2,3], 2) -> [[1,2],[1,3], [2,3]]
// Comb([1,2,3], 3) -> [[1,2,3]]
//
// This is basically all Edges that needs to be connected to new node n,
// each time creating a new graph.
type R struct {
	F [][]int
}

func (r *R) Add(x List) {
	r.F = append(r.F, x.Cp())
}

// List is just array of int but we can have better object manipulations.
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

// Cp copies l into a new list.
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

type LevelItem struct {
	Level int
	G     *G
}

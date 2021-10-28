package graph

import (
	"fmt"
	"math"
	"sort"
)

// G is a graph represented by a list of edges.
type G struct {
	Es []E // List of edges.
}

// Nodes returns the unique set of nodes [1..n]
func (g *G) Nodes() []int {
	nodes := map[int]bool{}
	for _, e := range g.Es {
		nodes[e.A] = true
		nodes[e.B] = true
	}

	res := []int{}
	for k := range nodes {
		res = append(res, k)
	}

	return res
}

// Cp copies the Graph G, to avoid issues with pass by values/reference.
func (g *G) Cp() *G {
	return &G{
		Es: g.Es,
	}
}

// String is a string representation of the graph G.
func (g *G) String() string {
	edges := []string{}
	for _, e := range g.Es {
		edges = append(edges, e.String())
	}
	sort.Slice(edges, func(i, j int) bool {
		return edges[i] <= edges[j]
	})
	return fmt.Sprintf("G(%v) = %v\n", len(g.Nodes()), edges)
}

// E is an undirected Graph Edge.
type E struct {
	A int
	B int
}

func (e *E) String() string {
	return fmt.Sprintf("(%v,%v)", int(math.Min(float64(e.A), float64(e.B))), int(math.Max(float64(e.A), float64(e.B))))
}

package main

import (
	"fmt"

	"com.github.anicolaspp/graphs/graph"
)

func main() {
	fmt.Println("Hello...")

	// f := &graph.R{F: map[string][]int{}}
	// graph.Comb([]int{1, 2}, 1, []int{}, f)

	// for i := 1; i <= 3; i++ {
	// 	f := &graph.R{F: [][]int{}}
	// 	graph.Comb([]int{1, 2, 3}, i, []int{}, f)
	// 	fmt.Println(f)
	// }

	fmt.Println("=======")

	gs := graph.Gen(3)

	fmt.Println("=======")
	fmt.Println(len(gs))
	for _, g := range gs {
		fmt.Println(g)
	}
}

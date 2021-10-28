package main

import (
	"fmt"

	"com.github.anicolaspp/graphs/graph"
)

func main() {
	fmt.Println("Hello...")

	fmt.Println("Running=======")

	gs := graph.Gen(6)

	fmt.Println("=======")
	fmt.Println(len(gs))
	for _, g := range gs {
		fmt.Println(g)
	}
}

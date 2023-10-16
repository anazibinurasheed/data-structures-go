package main

import "fmt"

// Graph represents an adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a Vertex to the Graph

func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		return
	}
	g.vertices = append(g.vertices, &Vertex{key: k})
}

// Add Edge

// getVertex

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacency list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("vertex :%v ->")

		for _, e := range v.adjacent {
			fmt.Printf("%v ->", e)
		}
		fmt.Println()
	}
}

func main() {

	g := &Graph{}

	for i := 0; i < 5; i++ {
		g.AddVertex(i)
	}

	fmt.Println(g)

	g.Print()
}

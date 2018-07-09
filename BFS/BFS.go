package main

import (
	"fmt"
)

type edge struct {
	source int
	dest   int
}

type graph struct {
	edges []edge
}

func (g *graph) getchildren(i int) []edge {
	children := []edge{}
	for _, e := range g.edges {
		if e.source == i {
			children = append(children, e)
		}
	}

	return children
}

func (g *graph) BFS(start, end int) int {
	queue := []int{}
	visited := [5]int{0}

	steps := 1
	index := 0

	if start == end {
		return 0
	}

	// append start node and mark as visited
	queue = append(queue, start)
	visited[start-1] = 1

	// while something is in the queue
	for index < len(queue) {
		c := g.getchildren(queue[index])
		visited[queue[index]-1] = 1

		// if any children
		if len(c) > 0 {
			for _, i := range c {
				if i.dest == end {
					return steps
				}

				// add to queue if not yet visited
				if visited[i.dest-1] == 0 {
					queue = append(queue, i.dest)
				}
			}

			steps++
		}

		index++

	}

	return -1
}

func (g *graph) DFS(start, end int) int {
	stack := []int{}
	visited := [5]int{0}

	top := 0
	bottom := 0

	stack = append(stack, start)
	visited[start-1] = 1
	top++
	j := 5
	fmt.Println("top is pointing to:", top)
	fmt.Println("bottom is pointing to:", bottom)
	for top > -1 {
		c := g.getchildren(stack[top-1])
		fmt.Printf("children for %d: %v\n", stack[top-1], c)
		visited[stack[top-1]-1] = 1

		bottom++
		top--

		if len(c) > 0 {
			for _, i := range c {
				if visited[i.dest-1] == 0 {
					stack = append(stack, i.dest)
					top++
				}
			}
		}

		fmt.Printf("visited: %v\n", visited)
		fmt.Printf("stack: %v\n", stack)
		j--
		fmt.Println("top is pointing to:", top)
		fmt.Println("bottom is pointing to:", bottom)
		if j == 0 {
			return 1
		}
	}

	return 1
}

func main() {

	edges := make([]edge, 5)
	edges[0] = edge{1, 2}
	edges[1] = edge{2, 3}
	edges[2] = edge{2, 4}
	edges[3] = edge{1, 5}
	edges[4] = edge{4, 6}

	graph := graph{edges}
	hasPath := graph.BFS(1, 3)
	fmt.Println(hasPath)

	hasPath = graph.DFS(1, 3)
	fmt.Println(hasPath)
}

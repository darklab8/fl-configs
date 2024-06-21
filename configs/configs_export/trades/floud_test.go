package trades

import (
	"fmt"
	"math"
	"testing"
)

func TestFloyd(t *testing.T) {
	fmt.Println(math.MaxFloat32)

	graph := NewGameGraph()
	// floyd := NewFloyder()
	graph.SetEdge("a", "b", 5)
	graph.SetEdge("a", "d", 10)
	graph.SetEdge("b", "c", 3)
	graph.SetEdge("c", "d", 1)

	floyd := NewFloyder(graph)
	floyd.Calculate()
	dist := floyd.dist

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if dist[i][j] == FloydMax {
				fmt.Printf("%7s", "INF")
			} else {
				fmt.Printf("%7d", dist[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("a -> c = ", GetDist(graph, dist, "a", "c"))
	fmt.Println("a -> b = ", GetDist(graph, dist, "a", "b"))
	fmt.Println("a -> b = ", GetDist(graph, dist, "a", "d"))
}

package trades

import (
	"fmt"
	"math"
	"testing"
)

func TestFloyd(t *testing.T) {
	fmt.Println(math.MaxFloat32)

	graph := NewFreelancerGraph()
	// floyd := NewFloyder()
	graph.SetEdge("a", "b", 5)
	graph.SetEdge("a", "d", 10)
	graph.SetEdge("b", "c", 3)
	graph.SetEdge("c", "d", 1)

	floyd := NewFloyder(graph)
	floyd.Calculate()

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if floyd.dist[i][j] == INF {
				fmt.Printf("%7s", "INF")
			} else {
				fmt.Printf("%7.0f", floyd.dist[i][j])
			}
		}
		fmt.Println()
	}

	fmt.Println("a", " -> ", "c", " = ", floyd.GetDist("a", "c"))
}

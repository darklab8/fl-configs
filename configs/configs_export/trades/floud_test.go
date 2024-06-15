package trades

import (
	"fmt"
	"testing"
)

func TestFloyd(t *testing.T) {
	floyd := NewFloyder()
	floyd.SetEdge("a", "b", 5)
	floyd.SetEdge("a", "d", 10)
	floyd.SetEdge("b", "c", 3)
	floyd.SetEdge("c", "d", 1)
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

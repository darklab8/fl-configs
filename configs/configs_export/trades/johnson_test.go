package trades

import (
	"fmt"
	"math"
	"testing"
)

// // Driver Code
func TestJohnson(t *testing.T) {
	var vertices int = 4
	var matrix [][]int = [][]int{
		{0, 0, -2, 0},
		{4, 0, 3, 0},
		{0, 0, 0, 2},
		{0, -1, 0, 0},
	}

	// Initialization
	var graph *Graph = NewGraphFromMatrix(vertices, matrix)

	// Function Call
	var distances [][]int = graph.johnsons()

	if distances == nil {
		fmt.Println("Negative weight cycle detected.")
		return
	}

	// The code fragment below outputs
	// an formatted distance matrix.
	// Its first row and first
	// column represent vertices
	fmt.Println("Distance matrix:")

	fmt.Printf("   \t")
	for i := 0; i < vertices; i++ {
		fmt.Printf("%3d\t", i)
	}

	for i := 0; i < vertices; i++ {
		fmt.Println()
		fmt.Printf("%3d\t", i)
		for j := 0; j < vertices; j++ {
			if distances[i][j] == math.MaxInt {
				fmt.Printf(" X\t")
			} else {
				fmt.Printf("%3d\t",
					distances[i][j])
			}
		}
	}
}

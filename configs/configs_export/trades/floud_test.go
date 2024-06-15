package trades

import (
	"fmt"
	"math"
	"testing"
)

const INF = math.MaxInt32

type VertexTest string

func (f *Floyder) floydWarshall() {
	for k := 0; k < 4; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				f.dist[i][j] = float64(math.Min(float64(f.dist[i][j]), float64(f.dist[i][k]+f.dist[k][j])))
			}
		}
	}
}

type Floyder struct {
	vertexes          []VertexTest
	dist              [][]float64
	index_by_nickname map[VertexTest]int
}

func NewFloyder(vertexes []VertexTest) *Floyder {
	f := &Floyder{
		vertexes:          vertexes,
		index_by_nickname: map[VertexTest]int{},
		dist:              make([][]float64, len(vertexes)),
	}
	// len_vertexes := len(vertexes)
	// for i := 0; i < len_vertexes; i++ {
	// 	f.dist[i] = make([]float64, len_vertexes)
	// 	for j := 0; j < len_vertexes; j++ {
	// 		f.dist[i][j] = INF
	// 	}
	// }
	return f
}

func (f *Floyder) Calculate() *Floyder {
	for index, vertex := range f.vertexes {
		f.index_by_nickname[vertex] = index
	}

	graph := [][]float64{
		{0, 5, INF, 10},
		{5, 0, 3, INF},
		{INF, 3, 0, 1},
		{10, INF, 1, 0},
	}

	f.dist = graph
	f.floydWarshall()

	return f
}

func (f *Floyder) GetDist(keya string, keyb string) float64 {
	return f.dist[f.index_by_nickname[VertexTest(keya)]][f.index_by_nickname[VertexTest(keyb)]]
}

func TestFloyd(t *testing.T) {
	var vertexes []VertexTest = []VertexTest{"a", "b", "c", "d"}

	floyd := NewFloyder(vertexes)
	// floyd.SetEdge()

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

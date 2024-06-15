package trades

import (
	"fmt"
	"math"
	"testing"
)

const INF = math.MaxInt32

type VertexTest string

/*
floydWarshall algirthm in a no matter directional graph
*/
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

	for index, vertex := range f.vertexes {
		f.index_by_nickname[vertex] = index
	}

	len_vertexes := len(vertexes)
	for i := 0; i < len_vertexes; i++ {
		f.dist[i] = make([]float64, len_vertexes)
		for j := 0; j < len_vertexes; j++ {
			f.dist[i][j] = INF
		}
	}

	for i := 0; i < len_vertexes; i++ {
		f.dist[i][i] = 0
	}
	return f
}

func (f *Floyder) Calculate() *Floyder {
	for k := 0; k < 4; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				f.dist[i][j] = float64(math.Min(float64(f.dist[i][j]), float64(f.dist[i][k]+f.dist[k][j])))
			}
		}
	}
	return f
}

func (f *Floyder) SetEdge(keya string, keyb string, distance float64) {
	f.dist[f.index_by_nickname[VertexTest(keya)]][f.index_by_nickname[VertexTest(keyb)]] = distance
	f.dist[f.index_by_nickname[VertexTest(keyb)]][f.index_by_nickname[VertexTest(keya)]] = distance
}

func (f *Floyder) GetDist(keya string, keyb string) float64 {
	return f.dist[f.index_by_nickname[VertexTest(keya)]][f.index_by_nickname[VertexTest(keyb)]]
}

func TestFloyd(t *testing.T) {
	var vertexes []VertexTest = []VertexTest{"a", "b", "c", "d"}

	floyd := NewFloyder(vertexes)
	floyd.SetEdge("a", "b", 5)
	floyd.SetEdge("a", "d", 10)
	floyd.SetEdge("b", "c", 3)
	floyd.SetEdge("c", "d", 1)

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
	fmt.Println("-------------")

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

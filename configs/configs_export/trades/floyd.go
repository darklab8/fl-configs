package trades

import (
	"math"
)

const INF = math.MaxInt32

type VertexTest string

/*
floydWarshall algirthm in a no matter directional graph
*/
type Floyder struct {
	matrix            map[VertexTest]map[VertexTest]float64
	dist              [][]float64
	index_by_nickname map[VertexTest]int
}

func (f *Floyder) SetEdge(keya string, keyb string, distance float64) {
	vertex, vertex_exists := f.matrix[VertexTest(keya)]
	if !vertex_exists {
		vertex = make(map[VertexTest]float64)
		f.matrix[VertexTest(keya)] = vertex
	}

	if _, vert_target_exists := f.matrix[VertexTest(keyb)]; !vert_target_exists {
		f.matrix[VertexTest(keyb)] = make(map[VertexTest]float64)
	}
	vertex[VertexTest(keyb)] = distance
}

func (f *Floyder) mapMatrixEdgeToFloyd(keya VertexTest, keyb VertexTest, distance float64) {
	f.dist[f.index_by_nickname[keya]][f.index_by_nickname[keyb]] = distance
	f.dist[f.index_by_nickname[keyb]][f.index_by_nickname[keya]] = distance
}

func NewFloyder() *Floyder {
	f := &Floyder{
		matrix:            make(map[VertexTest]map[VertexTest]float64),
		index_by_nickname: map[VertexTest]int{},
	}

	return f
}

func (f *Floyder) Calculate() *Floyder {

	len_vertexes := len(f.matrix)

	f.dist = make([][]float64, len_vertexes)

	for i := 0; i < len_vertexes; i++ {
		f.dist[i] = make([]float64, len_vertexes)
		for j := 0; j < len_vertexes; j++ {
			f.dist[i][j] = INF
		}
	}
	for i := 0; i < len_vertexes; i++ {
		f.dist[i][i] = 0
	}

	index := 0
	for vertex, _ := range f.matrix {
		f.index_by_nickname[vertex] = index
		index++
	}

	for vertex_source, vertex_targets := range f.matrix {
		for vertex_target_name, vertex_target_dist := range vertex_targets {
			f.mapMatrixEdgeToFloyd(vertex_source, vertex_target_name, vertex_target_dist)
		}
	}

	// optionally print for debugging
	// for i := 0; i < 4; i++ {
	// 	for j := 0; j < 4; j++ {
	// 		if floyd.dist[i][j] == INF {
	// 			fmt.Printf("%7s", "INF")
	// 		} else {
	// 			fmt.Printf("%7.0f", floyd.dist[i][j])
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println("-------------")

	for k := 0; k < 4; k++ {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				f.dist[i][j] = float64(math.Min(float64(f.dist[i][j]), float64(f.dist[i][k]+f.dist[k][j])))
			}
		}
	}
	return f
}

func (f *Floyder) GetDist(keya string, keyb string) float64 {
	return f.dist[f.index_by_nickname[VertexTest(keya)]][f.index_by_nickname[VertexTest(keyb)]]
}

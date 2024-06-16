package trades

import (
	"fmt"
	"math"
)

const INF = math.MaxFloat32

type VertexName string

type FreelancerGraph struct {
	matrix            map[VertexName]map[VertexName]float64
	index_by_nickname map[VertexName]int
}

func NewFreelancerGraph() *FreelancerGraph {
	return &FreelancerGraph{
		matrix:            make(map[VertexName]map[VertexName]float64),
		index_by_nickname: map[VertexName]int{},
	}
}

func (f *FreelancerGraph) SetEdge(keya string, keyb string, distance float64) {
	vertex, vertex_exists := f.matrix[VertexName(keya)]
	if !vertex_exists {
		vertex = make(map[VertexName]float64)
		f.matrix[VertexName(keya)] = vertex
	}

	if _, vert_target_exists := f.matrix[VertexName(keyb)]; !vert_target_exists {
		f.matrix[VertexName(keyb)] = make(map[VertexName]float64)
	}
	vertex[VertexName(keyb)] = distance
}

/*
floydWarshall algirthm in a no matter directional graph
*/
type Floyder struct {
	*FreelancerGraph
	dist [][]float64
}

func (f *Floyder) mapMatrixEdgeToFloyd(keya VertexName, keyb VertexName, distance float64) {
	f.dist[f.index_by_nickname[keya]][f.index_by_nickname[keyb]] = distance
	f.dist[f.index_by_nickname[keyb]][f.index_by_nickname[keya]] = distance
}

func NewFloyder(graph *FreelancerGraph) *Floyder {
	f := &Floyder{FreelancerGraph: graph}
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

	for k := 0; k < len_vertexes; k++ {
		fmt.Println("starting, k=", k)
		for i := 0; i < len_vertexes; i++ {
			for j := 0; j < len_vertexes; j++ {
				f.dist[i][j] = float64(math.Min(float64(f.dist[i][j]), float64(f.dist[i][k]+f.dist[k][j])))
			}
		}
	}
	return f
}

func (f *Floyder) GetDist(keya string, keyb string) float64 {
	return f.dist[f.index_by_nickname[VertexName(keya)]][f.index_by_nickname[VertexName(keyb)]]
}

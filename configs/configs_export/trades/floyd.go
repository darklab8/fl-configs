package trades

import (
	"fmt"
	"math"
)

const INF = math.MaxFloat32

type VertexName string

type GameGraph struct {
	matrix                        map[VertexName]map[VertexName]float64
	index_by_nickname             map[VertexName]int
	vertex_to_calculate_paths_for map[VertexName]bool
}

func NewGameGraph() *GameGraph {
	return &GameGraph{
		matrix:                        make(map[VertexName]map[VertexName]float64),
		index_by_nickname:             map[VertexName]int{},
		vertex_to_calculate_paths_for: make(map[VertexName]bool),
	}
}

func (f *GameGraph) SetEdge(keya string, keyb string, distance float64) {
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
	*GameGraph
	dist [][]int
}

var FloydMax = int(math.MaxInt / 2)

func (f *Floyder) mapMatrixEdgeToFloyd(keya VertexName, keyb VertexName, distance int) {
	f.dist[f.index_by_nickname[keya]][f.index_by_nickname[keyb]] = distance
	f.dist[f.index_by_nickname[keyb]][f.index_by_nickname[keya]] = distance
}

func NewFloyder(graph *GameGraph) *Floyder {
	f := &Floyder{GameGraph: graph}
	return f
}

func (f *Floyder) Calculate() *Floyder {

	len_vertexes := len(f.matrix)

	f.dist = make([][]int, len_vertexes)

	for i := 0; i < len_vertexes; i++ {
		f.dist[i] = make([]int, len_vertexes)
		for j := 0; j < len_vertexes; j++ {
			f.dist[i][j] = FloydMax
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
			f.mapMatrixEdgeToFloyd(vertex_source, vertex_target_name, int(vertex_target_dist))
		}
	}

	for k := 0; k < len_vertexes; k++ {
		if k%100 == 0 {
			fmt.Println("starting, k=", k)
		}
		for i := 0; i < len_vertexes; i++ {
			for j := 0; j < len_vertexes; j++ {
				if f.dist[i][k]+f.dist[k][j] < f.dist[i][j] {
					f.dist[i][j] = f.dist[i][k] + f.dist[k][j]
				}
			}
		}
	}
	return f
}

func GetDist[T any](f *GameGraph, dist [][]T, keya string, keyb string) T {
	return dist[f.index_by_nickname[VertexName(keya)]][f.index_by_nickname[VertexName(keyb)]]
}

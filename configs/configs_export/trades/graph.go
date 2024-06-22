package trades

/*
Game graph simplifies for us conversion of data from Freelancer space simulator to different graph algorithms.
*/

import (
	"math"
	"reflect"
)

const INF = math.MaxFloat32

type VertexName string

type GameGraph struct {
	matrix                        map[VertexName]map[VertexName]float64
	index_by_nickname             map[VertexName]int
	nickname_by_index             map[int]VertexName
	vertex_to_calculate_paths_for map[VertexName]bool
}

func NewGameGraph() *GameGraph {
	return &GameGraph{
		matrix:                        make(map[VertexName]map[VertexName]float64),
		index_by_nickname:             map[VertexName]int{},
		nickname_by_index:             make(map[int]VertexName),
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

func GetDist[T any](f *GameGraph, dist [][]T, keya string, keyb string) T {
	return dist[f.index_by_nickname[VertexName(keya)]][f.index_by_nickname[VertexName(keyb)]]
}

func GetPath(graph *GameGraph, parents [][]int, source_key string, target_key string) []string {
	// fmt.Println("get_path", source_key, target_key)
	S := []string{}
	u := graph.index_by_nickname[VertexName(target_key)] // target
	source := graph.index_by_nickname[VertexName(source_key)]

	if parents[source][u] != NO_PARENT || u == source {
		for {
			nickname := graph.nickname_by_index[u]
			S = append(S, string(nickname))
			u = parents[source][u]
			if u == NO_PARENT {
				break
			}
		}
	}
	ReverseSlice(S)
	return S
}

// panic if s is not a slice
func ReverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

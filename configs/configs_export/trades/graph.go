package trades

/*
Game graph simplifies for us conversion of data from Freelancer space simulator to different graph algorithms.
*/

import (
	"math"
	"reflect"
	"strings"
)

type VertexName string

type GameGraph struct {
	matrix                        map[VertexName]map[VertexName]float64
	Index_by_nickname             map[VertexName]int
	nickname_by_index             map[int]VertexName
	vertex_to_calculate_paths_for map[VertexName]bool
	AvgCruiseSpeed                int
}

func NewGameGraph(avgCruiseSpeed int) *GameGraph {
	return &GameGraph{
		matrix:                        make(map[VertexName]map[VertexName]float64),
		Index_by_nickname:             map[VertexName]int{},
		nickname_by_index:             make(map[int]VertexName),
		vertex_to_calculate_paths_for: make(map[VertexName]bool),
		AvgCruiseSpeed:                avgCruiseSpeed,
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
	return dist[f.Index_by_nickname[VertexName(keya)]][f.Index_by_nickname[VertexName(keyb)]]
}

type Path struct {
	Node     int
	NextNode int
	Dist     int
}

func GetPath(graph *GameGraph, parents [][]int, dist [][]int, source_key string, target_key string) []Path {
	// fmt.Println("get_path", source_key, target_key)
	S := []Path{}
	u := graph.Index_by_nickname[VertexName(target_key)] // target
	source := graph.Index_by_nickname[VertexName(source_key)]

	add_node := func(u int) {
		path_to_add := Path{
			Node: u,
		}
		if len(S) > 0 {
			path_to_add.NextNode = S[len(S)-1].Node
		} else {
			path_to_add.NextNode = NO_PARENT
		}
		if path_to_add.Node != NO_PARENT && path_to_add.NextNode != NO_PARENT {
			path_to_add.Dist = dist[path_to_add.Node][path_to_add.NextNode]
		}

		if path_to_add.Dist == int(graph.GetDistForTime(JumpHoleDelaySec)) {
			S[len(S)-1].Dist += path_to_add.Dist
			return
		}

		S = append(S, path_to_add)
	}
	add_node(u)

	if parents[source][u] != NO_PARENT || u == source {
		for {
			u = parents[source][u]

			nickname := graph.nickname_by_index[u]
			if strings.Contains(string(nickname), "trade_lane") {
				continue
			}

			add_node(u)

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

type DetailedPath struct {
	PrevName    string
	NextName    string
	PrevNode    int
	NextNode    int
	Dist        int
	TimeMinutes int
	TimeSeconds int
}

func (graph *GameGraph) GetPaths(parents [][]int, dist [][]int, source_key string, target_key string) []DetailedPath {
	var detailed_paths []DetailedPath

	paths := GetPath(graph, parents, dist, source_key, target_key)
	for _, path := range paths {
		minutes := int(math.Floor(float64(graph.GetTimeForDist(float64(path.Dist))) / 60))
		detailed_path := DetailedPath{
			PrevName:    string(graph.nickname_by_index[path.Node]),
			NextName:    string(graph.nickname_by_index[path.NextNode]),
			PrevNode:    path.Node,
			NextNode:    path.NextNode,
			Dist:        path.Dist,
			TimeMinutes: int(minutes),
			TimeSeconds: graph.GetTimeForDist(float64(path.Dist)) - minutes*60,
		}

		detailed_paths = append(detailed_paths, detailed_path)
	}

	return detailed_paths
}

package trades

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/go-utils/utils/timeit"
	"github.com/stretchr/testify/assert"
)

func TestTradeRoutes(t *testing.T) {

	configs := configs_mapped.TestFixtureConfigs()
	graph := MapConfigsToFGraph(configs, WithFreighterPaths(false))

	edges_count := 0
	for _, edges := range graph.matrix {
		edges_count += len(edges)
	}
	fmt.Println("graph.vertixes=", len(graph.matrix), "edges_count=", edges_count)

	// for profiling only stuff.
	f, err := os.Create("dijkstra_apsp.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	timeit.NewTimerF(func(m *timeit.Timer) {
		johnson := NewDijkstraApspFromGraph(graph)
		dist, parents := johnson.DijkstraApsp()

		// This version lf algorithm can provide you with distances only originating from space bases (and not proxy bases)
		// The rest of starting points were excluded for performance reasons
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li01_to_li02")=`, GetDist(graph, dist, "li01_01_base", "li01_to_li02"))
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li02_to_li01")=`, GetDist(graph, dist, "li01_01_base", "li02_to_li01"))
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li12_02_base")=`, GetDist(graph, dist, "li01_01_base", "li12_02_base"))
		dist1 := GetDist(graph, dist, "li01_01_base", "li01_02_base")
		dist2 := GetDist(graph, dist, "li01_01_base", "br01_01_base")
		dist3 := GetDist(graph, dist, "li01_01_base", "li12_02_base")
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li01_02_base")`, dist1)
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "br01_01_base")`, dist2)
		fmt.Println(`GetDist(graph, dist, "li01_01_base", "li12_02_base")`, dist3)
		assert.Greater(t, dist1, 0)
		assert.Greater(t, dist2, 0)
		assert.Greater(t, dist3, 0)

		fmt.Println("li01_01_base->br01_01_base")
		dist_ := GetDist(graph, dist, "li01_01_base", "br01_01_base")
		fmt.Println("dist=", dist_)
		fmt.Println("rough time=", float64(dist_)/float64(AvgCruiseSpeed))
		fmt.Println("li01_01_base->br01_01_base path:", GetPath(graph, parents, "li01_01_base", "br01_01_base"))
	}, timeit.WithMsg("trade routes calculated"))
}

package trades

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestFloyder(t *testing.T) {
	if true {
		// Takes too long time to calculate.
		// We should use Johnson (or use parallel Floyd)
		return
	}

	configs := configs_mapped.TestFixtureConfigs()
	graph := MapConfigsToFloyder(configs)

	floyd := NewFloyder(graph)
	floyd.Calculate()

	// call floyder.GetDist("li01_01_base", "li01_to_li02")
	// call floyder.GetDist("li01_to_li02", "li02_to_li01")
	// call floyder.GetDist("li02_to_li01", "li12_02_base")

	dist1 := GetDist(graph, floyd.dist, "li01_01_base", "li01_02_base")
	assert.Greater(t, dist1, float64(0))

	dist2 := GetDist(graph, floyd.dist, "li01_01_base", "br01_01_base")
	assert.Greater(t, dist2, float64(0))

	dist3 := GetDist(graph, floyd.dist, "li01_01_base", "li12_02_base")

	assert.Greater(t, dist3, float64(0))
}

// func TestJohnsoner(t *testing.T) {
// 	if true {
// 		// Takes too long time to calculate.
// 		// We should use Johnson (or use parallel Floyd)
// 		return
// 	}

// 	configs := configs_mapped.TestFixtureConfigs()
// 	floyder := MapConfigsToFloyder(configs)

// 	floyder.Calculate()

// 	// call floyder.GetDist("li01_01_base", "li01_to_li02")
// 	// call floyder.GetDist("li01_to_li02", "li02_to_li01")
// 	// call floyder.GetDist("li02_to_li01", "li12_02_base")

// 	dist1 := floyder.GetDist("li01_01_base", "li01_02_base")
// 	assert.Greater(t, dist1, float64(0))

// 	dist2 := floyder.GetDist("li01_01_base", "br01_01_base")
// 	assert.Greater(t, dist2, float64(0))

// 	dist3 := floyder.GetDist("li01_01_base", "li12_02_base")

// 	assert.Greater(t, dist3, float64(0))
// }

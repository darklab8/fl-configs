package trades

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestFloyder(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	floyder := MapConfigsToFloyder(configs)
	floyder.Calculate()

	dist1 := floyder.GetDist("li01_01_base", "li01_02_base")
	assert.Greater(t, dist1, float64(0))

	dist2 := floyder.GetDist("li01_01_base", "br01_01_base") // 2147483647 // distance should be lesser with tradelanes
	assert.Greater(t, dist2, float64(0))

}

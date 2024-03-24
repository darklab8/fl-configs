package configs_mapped

import (
	"testing"

	"github.com/darklab8/go-utils/goutils/utils"
)

func TestSimple(t *testing.T) {
	utils.TimeMeasure(func(m *utils.TimeMeasurer) {
		configs := TestFixtureConfigs()
		configs.Write(IsDruRun(true))
	})
}

package configs_mapped

import (
	"testing"

	"github.com/darklab8/go-utils/goutils/utils/time_measure"
)

func TestSimple(t *testing.T) {
	time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
		configs := TestFixtureConfigs()
		configs.Write(IsDruRun(true))
	})
}

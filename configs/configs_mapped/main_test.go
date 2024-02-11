package configs_mapped

import (
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/utils"
)

func TestSimple(t *testing.T) {
	utils.TimeMeasure(func() {
		configs := TestFixtureConfigs()
		configs.Write(IsDruRun(true))
	}, "dry run time")
}

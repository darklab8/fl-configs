package configs_export

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
)

func TestGetEngines(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)
	ids := exporter.GetTractors()
	exporter.GetEngines(ids)
	// assert.Greater(t, len(items), 0) # vanilla can have zero
}

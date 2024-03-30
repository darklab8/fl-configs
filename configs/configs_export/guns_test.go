package configs_export

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
)

func TestGetGuns(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	guns := exporter.GetGuns()
	_ = guns
	// assert.Greater(t, len(guns), 0)
	// exporter.infocards_parser.Get()
}

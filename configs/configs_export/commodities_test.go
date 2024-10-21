package configs_export

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestExportCommodities(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	useful_bases := FilterToUserfulBases(exporter.GetBases())
	useful_bases_by_nick := make(map[string]*Base)
	for _, base := range useful_bases {
		useful_bases_by_nick[base.Nickname] = base
	}

	items := exporter.GetCommodities(useful_bases_by_nick)
	assert.Greater(t, len(items), 0)

	fmt.Println(items[0])
}

package configs_export

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestGetOres(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	useful_bases := FilterToUserfulBases(exporter.GetBases())
	useful_bases_by_nick := make(map[string]*Base)
	for _, base := range useful_bases {
		useful_bases_by_nick[base.Nickname] = base
	}

	commodities := exporter.GetCommodities(useful_bases_by_nick)
	mining_operations := exporter.GetOres(commodities)
	assert.Greater(t, len(mining_operations), 0)
	fmt.Println("len(mining_operations)=", len(mining_operations))
}

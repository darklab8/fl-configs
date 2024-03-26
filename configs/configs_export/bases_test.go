package configs_export

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestExportBases(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	bases := exporter.Bases(NoNameIncluded(false))
	assert.Greater(t, len(bases), 0)
	assert.NotEqual(t, bases[0].Nickname, bases[1].Nickname)

	found_goods := false
	for _, base := range bases {
		if len(base.MarketGoods) > 0 {
			found_goods = true
		}
	}
	assert.True(t, found_goods, "expected finding some goods")
}

func TestExportMarketGoods(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	goods := exporter.GetMarketGoods()
	assert.Greater(t, len(goods), 0)
}

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
}

func TestExportMarketGoods(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	goods := exporter.GetMarketGoods()
	assert.Greater(t, len(goods), 0)
	assert.NotEqual(t, goods[0].GoodNickname, goods[1].GoodNickname)
}

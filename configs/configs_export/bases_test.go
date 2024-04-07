package configs_export

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestExportBases(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	bases := exporter.GetBases()
	assert.Greater(t, len(bases.Bases), 0)
	assert.NotEqual(t, bases.Bases[0].Nickname, bases.Bases[1].Nickname)

	found_goods := false
	for _, base := range bases.Bases {
		if len(base.MarketGoods) > 0 {
			found_goods = true
		}
	}
	assert.True(t, found_goods, "expected finding some goods")

	infocards := exporter.infocards_parser.Get()
	for _, base := range bases.Bases {
		if base.Nickname == "Br01_01_Base" {
			lines := infocards[base.Infocard]
			fmt.Println(base.Nickname, lines.Lines)
			assert.Greater(t, len(lines.Lines), 0, "expected finding lines in infocard")
			break
		}
	}
}

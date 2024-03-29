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

	bases := exporter.GetBases(NoNameIncluded(false))
	assert.Greater(t, len(bases), 0)
	assert.NotEqual(t, bases[0].Nickname, bases[1].Nickname)

	found_goods := false
	for _, base := range bases {
		if len(base.MarketGoods) > 0 {
			found_goods = true
		}
	}
	assert.True(t, found_goods, "expected finding some goods")

	infocards := exporter.infocards_parser.Get()
	for _, base := range bases {
		if base.Nickname == "Br01_01_Base" {
			lines := infocards.MapGet(base.Infocard)
			fmt.Println(base.Nickname, lines.Lines)
			assert.Greater(t, len(lines.Lines), 0, "expected finding lines in infocard")
			break
		}
	}
}

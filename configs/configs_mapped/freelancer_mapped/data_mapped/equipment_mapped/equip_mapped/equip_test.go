package equip_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/tests"
	"github.com/stretchr/testify/assert"
)

func TestReadSelectEquip(t *testing.T) {
	fileref := tests.FixtureFileFind().GetFile(FILENAME_SELECT_EQUIP)

	config := Read([]*iniload.IniLoader{iniload.NewLoader(fileref).Scan()})

	assert.Greater(t, len(config.Commodities), 0, "expected finding items")

	for _, commodity := range config.Commodities {
		commodity.IdsName.Get()
	}
}

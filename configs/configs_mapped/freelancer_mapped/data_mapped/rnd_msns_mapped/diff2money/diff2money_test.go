package diff2money

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	filesystem := filefind.FindConfigs(test_directory)
	fileref := filesystem.GetFile(FILENAME)

	loaded_market_ships := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(loaded_market_ships.DiffToMoney), 0, "expected finding some elements")
}

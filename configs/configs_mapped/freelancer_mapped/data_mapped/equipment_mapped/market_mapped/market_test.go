package market_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_filepath.Join(test_directory, FILENAME_SHIPS))

	loaded_market_ships := Read([]*iniload.IniLoader{iniload.NewLoader(fileref).Scan()})

	assert.Greater(t, len(loaded_market_ships.BaseGoods), 0, "market ships sections were not scanned")
}

func TestWriter(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	input_file := file.NewFile(utils_filepath.Join(test_directory, FILENAME_SHIPS))

	temp_directory := utils.GetCurrrentTempFolder()

	config := Read([]*iniload.IniLoader{iniload.NewLoader(input_file).Scan()})
	config.Files[0].SetOutputPath(utils_filepath.Join(temp_directory, FILENAME_SHIPS))
	config.Write()
}

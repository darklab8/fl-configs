package inireader

import (
	"testing"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_filepath.Join(test_directory, "market_ships.ini"))
	config := INIFile.Read(INIFile{}, fileref)

	assert.Greater(t, len(config.Sections), 0, "market ships sections were not scanned")
}

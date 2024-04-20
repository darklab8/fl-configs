/*
parse universe.ini
*/
package universe_mapped

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/settings/logus"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_types.FilePath(utils_filepath.Join(test_directory, FILENAME)))
	config := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(config.Bases), 0)
	assert.Greater(t, len(config.Systems), 0)
}

func TestReader2(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_types.FilePath(utils_filepath.Join(test_directory, "universe.vanilla.ini")))
	config := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(config.Bases), 0)
	assert.Greater(t, len(config.Systems), 0)
}

func TestIdentifySystemFiles(t *testing.T) {
	test_directory := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(test_directory)))
	filesystem := filefind.FindConfigs(freelancer_folder)
	logus.Log.Debug("filefind.FindConfigs" + fmt.Sprintf("%v", filesystem))

	universe_fileref := file.NewFile(utils_types.FilePath(filepath.Join(test_directory.ToString(), "testdata", FILENAME)))
	_ = Read(iniload.NewLoader(universe_fileref).Scan())
}

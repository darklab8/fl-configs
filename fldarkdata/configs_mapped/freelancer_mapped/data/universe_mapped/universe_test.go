/*
parse universe.ini
*/
package universe_mapped

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/configs_mapped/parserutils/filefind"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_types.FilePath(utils_filepath.Join(test_directory, FILENAME)))
	config := Config{}
	config.Read(fileref)

	assert.Greater(t, len(config.Bases), 0)
	assert.Greater(t, len(config.Systems), 0)
}

func TestIdentifySystemFiles(t *testing.T) {
	test_directory := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(test_directory)))
	filesystem := filefind.FindConfigs(freelancer_folder)
	logus.Log.Debug("filefind.FindConfigs" + fmt.Sprintf("%v", filesystem))

	config := Config{}
	universe_fileref := file.NewFile(utils_types.FilePath(filepath.Join(test_directory.ToString(), "testdata", FILENAME)))
	config.Read(universe_fileref)
}

package exe_mapped

import (
	"testing"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/stretchr/testify/assert"
)

func FixtureFLINIConfig() *Config {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_types.FilePath(utils_filepath.Join(test_directory, FILENAME_FL_INI)))
	config := (&Config{}).Read(fileref)
	return config
}

func TestReader(t *testing.T) {
	config := FixtureFLINIConfig()
	assert.Greater(t, len(config.Resources.Dll), 0)
}

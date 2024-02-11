package exe_mapped

import (
	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
)

func FixtureFLINIConfig() *Config {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_types.FilePath(utils_filepath.Join(test_directory, FILENAME_FL_INI)))
	config := (&Config{}).Read(fileref)
	return config
}

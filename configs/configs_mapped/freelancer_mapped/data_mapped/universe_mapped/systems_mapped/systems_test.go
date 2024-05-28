package systems_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(folder))))
	logus.Log.Debug("", utils_logus.FilePath(freelancer_folder))
	filesystem := filefind.FindConfigs(freelancer_folder)

	universe_ini := iniload.NewLoader(file.NewFile(filesystem.Hashmap[universe_mapped.FILENAME].GetFilepath())).Scan()
	universe_config := universe_mapped.Read(universe_ini)

	systems := Read(universe_config, filesystem)

	system, ok := systems.SystemsMap["br01"]
	assert.True(t, ok, "system should be present")

	system.Render()
}

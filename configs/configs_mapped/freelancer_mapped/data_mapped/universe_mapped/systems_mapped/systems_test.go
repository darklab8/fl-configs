package systems_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/settings/logger"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(folder))))
	logger.Log.Debug("", utils_logus.FilePath(freelancer_folder))
	filesystem := filefind.FindConfigs(freelancer_folder)

	universe_config := universe_mapped.Config{}
	universe_config.Read(file.NewFile(filesystem.Hashmap[universe_mapped.FILENAME].GetFilepath()))

	systems := (&Config{}).Read(&universe_config, filesystem)

	system, ok := systems.SystemsMap.MapGetValue("br01")
	assert.True(t, ok, "system should be present")

	_, ok = system.BasesByBase.MapGetValue("br01_01_base")
	assert.True(t, ok, "base should be present")

	system.Render()
}

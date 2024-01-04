package systems

import (
	"testing"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/freelancer/data/universe"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/parserutils/filefind"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/parserutils/filefind/file"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecycleParams(t *testing.T) {
	folder := utils.GetCurrentFolder()
	freelancer_folder := utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(utils_filepath.Dir(folder))))
	logus.Log.Debug("", logus_core.FilePath(freelancer_folder))
	filesystem := filefind.FindConfigs(freelancer_folder)

	universe_config := universe.Config{}
	universe_config.Read(file.NewFile(filesystem.Hashmap[universe.FILENAME].GetFilepath()))

	systems := (&Config{}).Read(&universe_config, filesystem)

	system, ok := systems.SystemsMap["br01"]
	assert.True(t, ok, "system should be present")

	_, ok = system.BasesByBase["br01_01_base"]
	assert.True(t, ok, "base should be present")

	system.Render()
}

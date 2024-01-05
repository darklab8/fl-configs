package configs_mapped

import (
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

func TestSimple(t *testing.T) {
	current_folder := utils.GetCurrentFolder()
	game_location := utils_filepath.Dir(utils_filepath.Dir(current_folder))

	parsed := NewMappedConfigs().Read(game_location)
	parsed.Write(IsDruRun(true))
}

package configs_mapped

import (
	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"
)

var parsed *MappedConfigs = nil

func TestFixtureConfigs() *MappedConfigs {
	if parsed != nil {
		return parsed
	}

	current_folder := utils.GetCurrentFolder()
	game_location := utils_filepath.Dir(utils_filepath.Dir(current_folder))

	parsed = NewMappedConfigs().Read(game_location)

	return parsed
}

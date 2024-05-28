package configs_mapped

import (
	"github.com/darklab8/fl-configs/configs/settings"
)

var parsed *MappedConfigs = nil

func TestFixtureConfigs() *MappedConfigs {
	if parsed != nil {
		return parsed
	}

	game_location := settings.GetGameLocation()
	parsed = NewMappedConfigs().Read(game_location)

	return parsed
}

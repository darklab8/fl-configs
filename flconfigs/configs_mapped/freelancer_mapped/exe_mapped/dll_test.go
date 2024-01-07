package exe_mapped

import (
	"testing"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/configs_fixtures"
	"github.com/stretchr/testify/assert"
)

func TestReadIncords(t *testing.T) {
	game_location := configs_fixtures.FixtureGameLocation()
	config := FixtureFLINIConfig()
	ids := GetAllInfocards(game_location, config.Resources.Dll)

	assert.Greater(t, len(ids), 0)
}

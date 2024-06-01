package npcranktodiff

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	filesystem := filefind.FindConfigs(configs_settings.GetGameLocation())
	fileref := filesystem.GetFile(FILENAME)

	loaded_market_ships := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(loaded_market_ships.NPCRankToDifficulties), 0, "expected finding some elements")
}

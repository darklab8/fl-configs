/*
parse universe.ini
*/
package universe_mapped

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	fileref := filefind.FindConfigs(configs_settings.GetGameLocation()).GetFile(FILENAME)
	config := Read(iniload.NewLoader(fileref).Scan())

	assert.Greater(t, len(config.Bases), 0)
	assert.Greater(t, len(config.Systems), 0)
}

func TestIdentifySystemFiles(t *testing.T) {

	filesystem := filefind.FindConfigs(configs_settings.GetGameLocation())
	logus.Log.Debug("filefind.FindConfigs" + fmt.Sprintf("%v", filesystem))

	universe_fileref := filesystem.GetFile(FILENAME)
	_ = Read(iniload.NewLoader(universe_fileref).Scan())
}

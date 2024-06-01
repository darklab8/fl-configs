package exe_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_settings"
)

func FixtureFLINIConfig() *Config {
	fileref := filefind.FindConfigs(configs_settings.GetGameLocation()).GetFile(FILENAME_FL_INI)
	config := Read(iniload.NewLoader(fileref).Scan())
	return config
}

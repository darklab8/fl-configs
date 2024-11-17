package solar_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/go-utils/utils/utils_types"
)

type Solar struct {
	semantic.Model
	Nickname       *semantic.String
	DockingSpheres []*semantic.String
}

type Config struct {
	*iniload.IniLoader
	Solars       []*Solar
	SolarsByNick map[string]*Solar
}

const (
	FILENAME utils_types.FilePath = "solararch.ini"
)

func Read(input_file *iniload.IniLoader) *Config {
	frelconfig := &Config{
		IniLoader:    input_file,
		SolarsByNick: make(map[string]*Solar),
	}

	for _, section := range input_file.SectionMap["[solar]"] {

		solar := &Solar{
			Nickname: semantic.NewString(section, "nickname", semantic.WithLowercaseS(), semantic.WithoutSpacesS()),
		}
		solar.Map(section)

		empathy_rate_key := "docking_sphere"
		for good_index, _ := range section.ParamMap[empathy_rate_key] {
			solar.DockingSpheres = append(solar.DockingSpheres,
				semantic.NewString(section, "docking_sphere", semantic.WithLowercaseS(), semantic.OptsS(semantic.Index(good_index)), semantic.WithoutSpacesS()))
		}

		frelconfig.Solars = append(frelconfig.Solars, solar)
		frelconfig.SolarsByNick[solar.Nickname.Get()] = solar

	}
	return frelconfig

}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

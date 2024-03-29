package ship_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-typelog/typelog"
)

type Ship struct {
	semantic.Model
	IdsName   *semantic.Int
	IdsInfo   *semantic.Int
	IdsInfo1  *semantic.Int
	IdsInfo2  *semantic.Int
	IdsInfo3  *semantic.Int
	Nickname  *semantic.String // matches value `ship` in goods
	Type      *semantic.String
	ShipClass *semantic.Int
}

type Config struct {
	Files []*iniload.IniLoader

	Ships    []*Ship
	ShipsMap *lower_map.KeyLoweredMap[string, *Ship]
}

func Read(files []*iniload.IniLoader) *Config {
	frelconfig := &Config{Files: files}
	frelconfig.Ships = make([]*Ship, 0, 100)
	frelconfig.ShipsMap = lower_map.NewKeyLoweredMap[string, *Ship]()

	for _, Iniconfig := range files {

		for _, section := range Iniconfig.SectionMap["[Ship]"] {
			ship := &Ship{}
			ship.Map(section)
			ship.Nickname = semantic.NewString(section, "nickname")
			ship.Type = semantic.NewString(section, "type")
			ship.ShipClass = semantic.NewInt(section, "ship_class")
			ship.IdsName = semantic.NewInt(section, "ids_name")
			ship.IdsInfo = semantic.NewInt(section, "ids_info")
			ship.IdsInfo1 = semantic.NewInt(section, "ids_info1")
			ship.IdsInfo2 = semantic.NewInt(section, "ids_info2")
			ship.IdsInfo3 = semantic.NewInt(section, "ids_info3")

			func() {
				defer func() {
					if r := recover(); r != nil {
						logus.Log.Debug("Recovered from grabbing ship IdsName. Error:\n", typelog.Any("recover", r))
					}
				}()

				ship.IdsName.Get()
				frelconfig.Ships = append(frelconfig.Ships, ship)
				frelconfig.ShipsMap.MapSet(ship.Nickname.Get(), ship)
			}()

		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() []*file.File {
	var files []*file.File
	for _, file := range frelconfig.Files {
		inifile := file.Render()
		inifile.Write(inifile.File)
		files = append(files, inifile.File)
	}
	return files
}

package const_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
)

const (
	FILENAME = "constants.ini"
)

type ShieldEquipConsts struct {
	semantic.Model
	HULL_DAMAGE_FACTOR *semantic.Float
}

type Config struct {
	*iniload.IniLoader

	ShieldEquipConsts *ShieldEquipConsts
}

func Read(input_file *iniload.IniLoader) *Config {
	frelconfig := &Config{
		IniLoader: input_file,
	}
	if groups, ok := frelconfig.SectionMap["[ShieldEquipConsts]"]; ok {
		shield_consts := &ShieldEquipConsts{}
		shield_consts.Map(groups[0])
		shield_consts.HULL_DAMAGE_FACTOR = semantic.NewFloat(groups[0], "hull_damage_factor", semantic.Precision(2))

		frelconfig.ShieldEquipConsts = shield_consts
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

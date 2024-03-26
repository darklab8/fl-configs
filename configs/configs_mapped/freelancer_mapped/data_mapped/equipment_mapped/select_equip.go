package equipment_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type CommodityEquip struct {
	semantic.Model

	Nickname          *semantic.String
	IdsName           *semantic.Int
	IdsInfo           *semantic.Int
	UnitsPerContainer *semantic.Int
	PodApperance      *semantic.String
	LootAppearance    *semantic.String
	DecayPerSecond    *semantic.Int
	Volume            *semantic.Int
	HitPts            *semantic.Int
}

type ConfigSelectEquip struct {
	semantic.ConfigModel

	Commodities []*CommodityEquip
}

const (
	FILENAME_SELECT_EQUIP utils_types.FilePath = "select_equip.ini"
)

func ReadSelectEquip(input_file *file.File) *ConfigSelectEquip {
	frelconfig := &ConfigSelectEquip{}
	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
	frelconfig.Commodities = make([]*CommodityEquip, 0, 100)

	for _, section := range iniconfig.SectionMap["[Commodity]"] {
		commodity := &CommodityEquip{}
		commodity.Map(section)
		commodity.Nickname = semantic.NewString(section, "nickname")
		commodity.IdsName = semantic.NewInt(section, "ids_name")
		commodity.IdsInfo = semantic.NewInt(section, "ids_info")
		commodity.UnitsPerContainer = semantic.NewInt(section, "units_per_container")
		commodity.PodApperance = semantic.NewString(section, "pod_appearance")
		commodity.LootAppearance = semantic.NewString(section, "loot_appearance")
		commodity.DecayPerSecond = semantic.NewInt(section, "decay_per_second")
		commodity.Volume = semantic.NewInt(section, "volume")
		commodity.HitPts = semantic.NewInt(section, "hit_pts")

		frelconfig.Commodities = append(frelconfig.Commodities, commodity)

	}
	return frelconfig
}

func (frelconfig *ConfigSelectEquip) Write() *file.File {

	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

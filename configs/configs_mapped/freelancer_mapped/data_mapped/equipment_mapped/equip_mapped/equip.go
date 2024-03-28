package equip_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Commodity struct {
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

type ConfigFile struct {
	semantic.ConfigModel
}

type Config struct {
	Files []*ConfigFile

	Commodities    []*Commodity
	CommoditiesMap *lower_map.KeyLoweredMap[string, *Commodity]
}

const (
	FILENAME_SELECT_EQUIP utils_types.FilePath = "select_equip.ini"
)

func Read(input_files []*file.File) *Config {
	frelconfig := &Config{}
	frelconfig.Commodities = make([]*Commodity, 0, 100)
	frelconfig.CommoditiesMap = lower_map.NewKeyLoweredMap[string, *Commodity]()

	for _, input_file := range input_files {
		fileconfig := &ConfigFile{}
		iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
		fileconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
		frelconfig.Files = append(frelconfig.Files, fileconfig)

		for _, section := range iniconfig.SectionMap["[Commodity]"] {
			commodity := &Commodity{}
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
			frelconfig.CommoditiesMap.MapSet(commodity.Nickname.Get(), commodity)
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

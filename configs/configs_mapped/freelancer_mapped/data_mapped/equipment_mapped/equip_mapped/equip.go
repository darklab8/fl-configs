package equip_mapped

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Item struct {
	semantic.Model

	Category string
	Nickname *semantic.String
	IdsName  *semantic.Int
	IdsInfo  *semantic.Int
}

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

type Config struct {
	Files []*iniload.IniLoader

	Commodities    []*Commodity
	CommoditiesMap *lower_map.KeyLoweredMap[string, *Commodity]

	Items    []*Item
	ItemsMap *lower_map.KeyLoweredMap[string, *Item]
}

const (
	FILENAME_SELECT_EQUIP utils_types.FilePath = "select_equip.ini"
)

func Read(files []*iniload.IniLoader) *Config {
	frelconfig := &Config{Files: files}
	frelconfig.Commodities = make([]*Commodity, 0, 100)
	frelconfig.CommoditiesMap = lower_map.NewKeyLoweredMap[string, *Commodity]()
	frelconfig.Items = make([]*Item, 0, 100)
	frelconfig.ItemsMap = lower_map.NewKeyLoweredMap[string, *Item]()

	for _, file := range files {
		for _, section := range file.SectionMap["[Commodity]"] {
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

		for _, section := range file.Sections {
			item := &Item{}
			item.Map(section)
			item.Category = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(string(section.Type), "[", ""), "]", ""))
			item.Nickname = semantic.NewString(section, "nickname", semantic.OptsS(semantic.Optional()))
			item.IdsName = semantic.NewInt(section, "ids_name", semantic.Optional())
			item.IdsInfo = semantic.NewInt(section, "ids_info", semantic.Optional())
			frelconfig.Items = append(frelconfig.Items, item)
			frelconfig.ItemsMap.MapSet(item.Nickname.Get(), item)
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

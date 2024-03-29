package market_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

// Not implemented. Create SemanticMultiKeyValue
type MarketGood struct {
	semantic.Model
	Nickname *semantic.String // 0

	LevelRequired *semantic.Int   // 1
	RepRequired   *semantic.Float // 2

	IsBuyOnly     *semantic.IntBool // 5
	PriceModifier *semantic.Float   // 6
}

type BaseGood struct {
	semantic.Model
	Base *semantic.String

	MarketGoods []*MarketGood
}

type ConfigFile struct {
	semantic.ConfigModel
}

type Config struct {
	Files []*ConfigFile

	BaseGoods []*BaseGood
}

const (
	FILENAME_SHIPS            utils_types.FilePath = "market_ships.ini"
	FILENAME_COMMODITIES      utils_types.FilePath = "market_commodities.ini"
	FILENAME_MISC             utils_types.FilePath = "market_misc.ini"
	BaseGoodType                                   = "[BaseGood]"
	KEY_MISSMATCH_SYSTEM_FILE                      = "missmatched_universe_system_and_file"
	KEY_MARKET_GOOD                                = "marketgood"
	KEY_BASE                                       = "base"
)

func Read(input_files []*file.File) *Config {

	frelconfig := &Config{}
	frelconfig.BaseGoods = make([]*BaseGood, 0)

	for _, input_file := range input_files {
		fileconfig := &ConfigFile{}
		iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
		fileconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
		frelconfig.Files = append(frelconfig.Files, fileconfig)

		for _, section := range iniconfig.Sections {
			base_to_add := &BaseGood{}
			base_to_add.Map(section)
			base_to_add.Base = semantic.NewString(section, KEY_BASE)

			for good_index, _ := range section.ParamMap[KEY_MARKET_GOOD] {
				good_to_add := &MarketGood{}
				good_to_add.Map(section)
				good_to_add.Nickname = semantic.NewString(section, KEY_MARKET_GOOD, semantic.OptsS(semantic.Index(good_index)))
				good_to_add.LevelRequired = semantic.NewInt(section, KEY_MARKET_GOOD, semantic.Index(good_index), semantic.Order(1))
				good_to_add.RepRequired = semantic.NewFloat(section, KEY_MARKET_GOOD, semantic.Precision(2), semantic.Index(good_index), semantic.Order(2))
				good_to_add.IsBuyOnly = semantic.NewIntBool(section, KEY_MARKET_GOOD, semantic.Index(good_index), semantic.Order(5))
				good_to_add.PriceModifier = semantic.NewFloat(section, KEY_MARKET_GOOD, semantic.Precision(2), semantic.Index(good_index), semantic.Order(6))
				base_to_add.MarketGoods = append(base_to_add.MarketGoods, good_to_add)
			}

			frelconfig.BaseGoods = append(frelconfig.BaseGoods, base_to_add)
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

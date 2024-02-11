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
	Name *semantic.String
	// Values SemanticIntArray
}

type BaseGood struct {
	semantic.Model
	Base *semantic.String
	// TODO Goods          *SemanticMultiKey[MarketGood] (GetAll)
}

type Config struct {
	semantic.ConfigModel

	BaseGoods []*BaseGood
	Comments  []string
}

const (
	FILENAME_SHIPS            utils_types.FilePath = "market_ships.ini"
	FILENAME_COMMODITIES      utils_types.FilePath = "market_commodities.ini"
	FILENAME_MISC             utils_types.FilePath = "market_misc.ini"
	BaseGoodType                                   = "[BaseGood]"
	KEY_NAME                                       = "name"
	KEY_RECYCLE                                    = "is_recycle_candidate"
	KEY_MISSMATCH_SYSTEM_FILE                      = "missmatched_universe_system_and_file"
	KEY_MARKET_GOOD                                = "marketgood"
	KEY_BASE                                       = "base"
)

func (frelconfig *Config) Read(input_file *file.File) *Config {
	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
	frelconfig.BaseGoods = make([]*BaseGood, 0)

	for _, section := range iniconfig.Sections {
		base_to_add := &BaseGood{}
		base_to_add.Map(section)
		base_to_add.Base = semantic.NewString(section, KEY_BASE, semantic.TypeVisible, inireader.REQUIRED_p)
		frelconfig.BaseGoods = append(frelconfig.BaseGoods, base_to_add)
	}
	frelconfig.Comments = iniconfig.Comments
	return frelconfig
}

func (frelconfig *Config) Write() *file.File {

	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

package equipment_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Commodity struct {
	semantic.Model
	Nickname  *semantic.String
	Equipment *semantic.String
	Category  *semantic.String

	Price         *semantic.Int
	Combinable    *semantic.String // should be StrBool
	GoodSellPrice *semantic.Float
	BadBuyPrice   *semantic.Float
	BadSellPrice  *semantic.Float
	GoodBuyPrice  *semantic.Float
	ShopArchetype *semantic.Path
	ItemIcon      *semantic.Path
	JumpDist      *semantic.Int
}

type Good struct {
	semantic.Model
	Category *semantic.String
	Nickname *semantic.String
	Price    *semantic.Int
}

type ConfigFile struct {
	semantic.ConfigModel
}

type Config struct {
	Files []*ConfigFile

	Goods          []*Good
	GoodsMap       *lower_map.KeyLoweredMap[string, *Good]
	Commodities    []*Commodity
	CommoditiesMap *lower_map.KeyLoweredMap[string, *Commodity]
}

const (
	FILENAME utils_types.FilePath = "goods.ini"
)

func Read(input_files []*file.File) *Config {
	frelconfig := &Config{}
	frelconfig.Commodities = make([]*Commodity, 0, 100)
	frelconfig.CommoditiesMap = lower_map.NewKeyLoweredMap[string, *Commodity]()

	frelconfig.Goods = make([]*Good, 0, 100)
	frelconfig.GoodsMap = lower_map.NewKeyLoweredMap[string, *Good]()

	for _, input_file := range input_files {
		fileconfig := &ConfigFile{}
		iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
		fileconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
		frelconfig.Files = append(frelconfig.Files, fileconfig)

		for _, section := range iniconfig.SectionMap["[Good]"] {
			good := &Good{}
			good.Map(section)
			good.Nickname = semantic.NewString(section, "nickname")
			good.Category = semantic.NewString(section, "category")
			good.Price = semantic.NewInt(section, "price", semantic.Optional())
			frelconfig.Goods = append(frelconfig.Goods, good)
			frelconfig.GoodsMap.MapSet(good.Nickname.Get(), good)

			category := good.Category.Get()
			switch category {
			case "commodity":
				commodity := &Commodity{}
				commodity.Map(section)
				commodity.Category = semantic.NewString(section, "category")
				commodity.Nickname = semantic.NewString(section, "nickname")
				commodity.Equipment = semantic.NewString(section, "equipment")
				commodity.Price = semantic.NewInt(section, "price")
				commodity.Combinable = semantic.NewString(section, "combinable")
				commodity.GoodSellPrice = semantic.NewFloat(section, "good_sell_price", semantic.Precision(2))
				commodity.BadBuyPrice = semantic.NewFloat(section, "bad_buy_price", semantic.Precision(2))
				commodity.BadSellPrice = semantic.NewFloat(section, "bad_sell_price", semantic.Precision(2))
				commodity.GoodBuyPrice = semantic.NewFloat(section, "good_buy_price", semantic.Precision(2))
				commodity.ShopArchetype = semantic.NewPath(section, "shop_archetype")
				commodity.ItemIcon = semantic.NewPath(section, "item_icon")
				commodity.JumpDist = semantic.NewInt(section, "jump_dist")

				frelconfig.Commodities = append(frelconfig.Commodities, commodity)
				frelconfig.CommoditiesMap.MapSet(commodity.Nickname.Get(), commodity)
			}

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

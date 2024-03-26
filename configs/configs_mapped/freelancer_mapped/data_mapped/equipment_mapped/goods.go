package equipment_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"

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

type Config struct {
	semantic.ConfigModel

	Commodities []*Commodity
}

const (
	FILENAME utils_types.FilePath = "goods.ini"
	// GOOD_KEY                      = "[Good]"
)

func Read(input_file *file.File) *Config {
	frelconfig := &Config{}
	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
	frelconfig.Commodities = make([]*Commodity, 0)

	for _, section := range iniconfig.Sections {
		commodity := &Commodity{}
		commodity.Map(section)
		commodity.Category = semantic.NewString(section, "category")

		switch category := commodity.Category.Get(); category {
		case "commodity":
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
		}
	}
	return frelconfig
}

func (frelconfig *Config) Write() *file.File {

	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

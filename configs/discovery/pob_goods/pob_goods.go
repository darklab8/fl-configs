package pob_goods

import (
	"encoding/base64"
	"encoding/json"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"
)

type ShopItem struct {
	Id        int `json:"id"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
	SellPrice int `json:"sell_price"`
	MinStock  int `json:"min_stock"`
	MaxStock  int `json:"max_stock"`
}

func (good ShopItem) BaseSells() bool {
	return good.Quantity > good.MinStock
}
func (good ShopItem) BaseBuys() bool {
	return good.Quantity < good.MaxStock
}

type Base struct {
	Name      string
	Nickname  string
	ShopItems []ShopItem `json:"shop_items"`

	SystemHash      *flhash.HashCode `json:"system"`      //: 2745655887,
	Pos             *string          `json:"pos"`         //: "299016, 33, -178",
	AffiliationHash *flhash.HashCode `json:"affiliation"` //: 2620,
	Level           *int             `json:"level"`       //: 1,
	Money           *int             `json:"money"`       //: 0,
	Health          *int             `json:"health"`      //: 50,
	DefenseMode     *int             `json:"defensemode"` //: 1,
}

type Config struct {
	BasesByName map[string]*Base `json:"bases"`
	Timestamp   string           `json:"timestamp"`
	Bases       []*Base
}

func Read(file *file.File) *Config {
	byteValue, err := file.ReadBytes()
	logus.Log.CheckFatal(err, "failed to read file")

	var conf *Config
	json.Unmarshal(byteValue, &conf)

	for base_name, base := range conf.BasesByName {
		base.Name = base_name

		hash := base64.StdEncoding.EncodeToString([]byte(base.Name))
		base.Nickname = hash
		conf.Bases = append(conf.Bases, base)
	}

	return conf
}

func (frelconfig *Config) Write() *file.File {
	return &file.File{}
}

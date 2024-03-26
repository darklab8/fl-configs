package configs_export

import (
	"math"
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
)

type GoodType string

const (
	TypeCommodity GoodType = "commodity"
)

type MarketGood struct {
	Name     string
	Nickname string
	Type     GoodType

	LevelRequired int
	RepRequired   float64

	IsBuyOnly     bool
	PriceModifier float64
	PriceBase     int
	Price         int
}

func (e *Exporter) GetMarketGoods() map[string][]MarketGood {
	var GoodsPerBase map[string][]MarketGood = make(map[string][]MarketGood)

	for _, base_good := range e.configs.MarketCommidities.BaseGoods {
		base_nickname := base_good.Base.Get()

		var MarketGoods []MarketGood = make([]MarketGood, 0, 20)
		for _, market_good := range base_good.MarketGoods {

			commodity_selequip := e.configs.SelectEquip.CommoditiesMap.MapGet(strings.ToLower(market_good.Nickname.Get()))

			var Name infocard.Infoname
			if infoname, ok := e.configs.Infocards.Infonames[commodity_selequip.IdsName.Get()]; ok {
				Name = infoname
			}

			commodity_good := e.configs.Goods.CommoditiesMap.MapGet(strings.ToLower(market_good.Nickname.Get()))

			MarketGoods = append(MarketGoods, MarketGood{
				Name:          string(Name),
				Nickname:      market_good.Nickname.Get(),
				Type:          TypeCommodity,
				LevelRequired: market_good.LevelRequired.Get(),
				RepRequired:   market_good.RepRequired.Get(),
				IsBuyOnly:     market_good.IsBuyOnly.Get(),
				PriceModifier: market_good.PriceModifier.Get(),
				PriceBase:     commodity_good.Price.Get(),
				Price:         int(math.Floor(float64(commodity_good.Price.Get()) * market_good.PriceModifier.Get())),
			})
		}

		GoodsPerBase[base_nickname] = MarketGoods

	}
	return GoodsPerBase
}

package configs_export

import (
	"math"

	"github.com/darklab8/fl-configs/configs/settings/logus"
)

type MarketGood struct {
	Name     string
	Nickname string
	Type     string

	LevelRequired int
	RepRequired   float64
	Infocard      Infocard

	IsBuyOnly     bool
	PriceModifier float64
	PriceBase     int
	Price         int
}

func NameWithSpacesOnly(word string) bool {
	for _, ch := range word {
		if ch != ' ' {
			return false
		}
	}
	return true
}

func (e *Exporter) getMarketGoods() map[string][]MarketGood {
	var goods_per_base map[string][]MarketGood = make(map[string][]MarketGood)

	for _, base_good := range e.configs.Market.BaseGoods {
		base_nickname := base_good.Base.Get()

		var MarketGoods []MarketGood = make([]MarketGood, 0, 200)
		for _, market_good := range base_good.MarketGoods {

			var nickname string = market_good.Nickname.Get()
			var price_base int
			var Name string
			var category string
			var infocard_res Infocard
			if good, found_good := e.configs.Goods.GoodsMap.MapGetValue(nickname); found_good {
				price_base = good.Price.Get()

				category = good.Category.Get()
				switch category {
				default:
					if equip, ok := e.configs.Equip.ItemsMap.MapGetValue(nickname); ok {
						if infoname, ok := e.configs.Infocards.Infonames[equip.IdsName.Get()]; ok {
							Name = string(infoname)
							category = equip.Category
						}

						if infocard, infocard_exists := e.configs.Infocards.Infocards[equip.IdsInfo.Get()]; infocard_exists {
							infocardStart, err := infocard.XmlToText()
							logus.Log.CheckError(err, "failed to xml infocard")
							infocard_res.Lines = append(infocard_res.Lines, infocardStart...)
						}

					}
				case "ship":
					ship := e.configs.Goods.ShipsMap.MapGet(good.Nickname.Get())

					ship_hull := e.configs.Goods.ShipHullsMap.MapGet(ship.Hull.Get())
					price_base = ship_hull.Price.Get()

					// Infocard data
					ship_nickname := ship_hull.Ship.Get()
					shiparch := e.configs.Shiparch.ShipsMap.MapGet(ship_nickname)

					if infoname, ok := e.configs.Infocards.Infonames[shiparch.IdsName.Get()]; ok {
						Name = string(infoname)
					}
					if infocard, infocard_exists := e.configs.Infocards.Infocards[shiparch.IdsInfo.Get()]; infocard_exists {
						infocardStart, err := infocard.XmlToText()
						logus.Log.CheckError(err, "failed to xml infocard")
						infocard_res.Lines = append(infocard_res.Lines, infocardStart...)
					}
					if infocard, infocard_exists := e.configs.Infocards.Infocards[shiparch.IdsInfo1.Get()]; infocard_exists {
						infocardStart, err := infocard.XmlToText()
						logus.Log.CheckError(err, "failed to xml infocard")
						infocard_res.Lines = append(infocard_res.Lines, infocardStart...)
					}
					if infocard, infocard_exists := e.configs.Infocards.Infocards[shiparch.IdsInfo2.Get()]; infocard_exists {
						infocardStart, err := infocard.XmlToText()
						logus.Log.CheckError(err, "failed to xml infocard")
						infocard_res.Lines = append(infocard_res.Lines, infocardStart...)
					}
					if infocard, infocard_exists := e.configs.Infocards.Infocards[shiparch.IdsInfo3.Get()]; infocard_exists {
						infocardStart, err := infocard.XmlToText()
						logus.Log.CheckError(err, "failed to xml infocard")
						infocard_res.Lines = append(infocard_res.Lines, infocardStart...)
					}
				}

			}

			if NameWithSpacesOnly(Name) {
				Name = ""
			}

			MarketGoods = append(MarketGoods, MarketGood{
				Name:          Name,
				Nickname:      nickname,
				Type:          category,
				LevelRequired: market_good.LevelRequired.Get(),
				RepRequired:   market_good.RepRequired.Get(),
				IsBuyOnly:     market_good.IsBuyOnly.Get(),
				PriceModifier: market_good.PriceModifier.Get(),
				PriceBase:     price_base,
				Price:         int(math.Floor(float64(price_base) * market_good.PriceModifier.Get())),
				Infocard:      infocard_res,
			})
		}

		goods_per_base[base_nickname] = append(goods_per_base[base_nickname], MarketGoods...)

	}
	return goods_per_base
}

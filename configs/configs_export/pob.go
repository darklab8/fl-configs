package configs_export

import (
	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
)

var pob_produced_cached map[string]bool

func (e *Exporter) pob_produced() map[string]bool {
	if pob_produced_cached != nil {
		return pob_produced_cached
	}

	pob_produced_cached = make(map[string]bool)

	if e.configs.Discovery == nil {
		return pob_produced_cached
	}

	for _, recipe := range e.configs.Discovery.BaseRecipeItems.Recipes {
		for _, produced := range recipe.ProcucedItem {
			pob_produced_cached[produced.Get()] = true
		}
	}
	return pob_produced_cached
}

const pob_crafts_nickname = "pob_crafts"

func (e *Exporter) EnhanceBasesWithPobCrafts(bases []*Base) []*Base {
	pob_produced := e.pob_produced()

	base := &Base{
		Name:               "PoB Crafts",
		MarketGoodsPerNick: make(map[CommodityKey]MarketGood),
		Nickname:           cfgtype.BaseUniNick(pob_crafts_nickname),
		Infocard:           InfocardKey(pob_crafts_nickname),
		SystemNickname:     "neverwhere",
		System:             "Neverwhere",
		Region:             "Neverwhere",
		FactionName:        "Player Base Crafts",
	}

	base.Archetypes = append(base.Archetypes, pob_crafts_nickname)

	for produced, _ := range pob_produced {
		market_good := MarketGood{
			Nickname:             produced,
			NicknameHash:         flhash.HashNickname(produced),
			Infocard:             InfocardKey(produced),
			BaseSells:            true,
			Type:                 "craftable",
			ShipClass:            -1,
			IsServerSideOverride: true,
		}
		e.Hashes[market_good.Nickname] = market_good.NicknameHash

		if good_ship, ok := e.configs.Goods.ShipsMap[produced]; ok {
			hull_name := good_ship.Hull.Get()
			if good_shiphull, ok := e.configs.Goods.ShipHullsMap[hull_name]; ok {
				ship_nick := good_shiphull.Ship.Get()
				if equipment, ok := e.configs.Shiparch.ShipsMap[ship_nick]; ok {
					market_good.Name = e.GetInfocardName(equipment.IdsName.Get(), market_good.Nickname)
				}
			}
		} else {
			if equip, ok := e.configs.Equip.ItemsMap[produced]; ok {
				market_good.Name = e.GetInfocardName(equip.IdsName.Get(), produced)
			}
		}
		market_good_key := GetCommodityKey(market_good.Nickname, market_good.ShipClass)
		base.MarketGoodsPerNick[market_good_key] = market_good
	}

	var sb []string
	sb = append(sb, base.Name)
	sb = append(sb, `This is only pseudo base to show availability of player base crafts`)

	e.Infocards[InfocardKey(base.Nickname)] = sb

	bases = append(bases, base)
	return bases
}

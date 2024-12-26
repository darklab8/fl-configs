package configs_export

import (
	"fmt"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
)

var findable_in_loot map[string]bool

func (e *Exporter) findable_in_loot() map[string]bool {
	if findable_in_loot != nil {
		return findable_in_loot
	}

	findable_in_loot = make(map[string]bool)

	for _, system := range e.configs.Systems.Systems {
		for _, wreck := range system.Wrecks {
			louadout_nickname := wreck.Loadout.Get()
			if loadout, ok := e.configs.Loadouts.LoadoutsByNick[louadout_nickname]; ok {
				for _, cargo := range loadout.Cargos {
					findable_in_loot[cargo.Nickname.Get()] = true
				}
			}
		}
	}

	for _, npc_arch := range e.configs.NpcShips.NpcShips {
		loadout_nickname := npc_arch.Loadout.Get()
		if loadout, ok := e.configs.Loadouts.LoadoutsByNick[loadout_nickname]; ok {
			for _, cargo := range loadout.Cargos {
				findable_in_loot[cargo.Nickname.Get()] = true
			}
		}
	}
	return findable_in_loot
}

/*
It fixes issue of Guns obtainable only via wrecks being invisible
*/
const (
	BaseLootableName     = "Lootable"
	BaseLootableFaction  = "Wrecks and Missions"
	BaseLootableNickname = "base_loots"
)

func (e *Exporter) EnhanceBasesWithLoot(bases []*Base) []*Base {

	in_wrecks := e.findable_in_loot()

	base := &Base{
		Name:               "Lootable",
		MarketGoodsPerNick: make(map[CommodityKey]MarketGood),
		Nickname:           cfgtype.BaseUniNick(BaseLootableNickname),
		Infocard:           InfocardKey(BaseLootableNickname),
		SystemNickname:     "neverwhere",
		System:             "Neverwhere",
		Region:             "Neverwhere",
		FactionName:        BaseLootableFaction,
	}

	base.Archetypes = append(base.Archetypes, BaseLootableNickname)

	for wreck, _ := range in_wrecks {
		market_good := MarketGood{
			Nickname:             wreck,
			NicknameHash:         flhash.HashNickname(wreck),
			Infocard:             InfocardKey(wreck),
			BaseSells:            true,
			Type:                 "lootable",
			ShipClass:            -1,
			IsServerSideOverride: true,
		}
		e.Hashes[market_good.Nickname] = market_good.NicknameHash

		if good, found_good := e.configs.Goods.GoodsMap[market_good.Nickname]; found_good {
			category := good.Category.Get()
			market_good.Type = fmt.Sprintf("%s loot", category)
			if equip, ok := e.configs.Equip.ItemsMap[market_good.Nickname]; ok {
				market_good.Type = fmt.Sprintf("%s loot", equip.Category)
				e.exportInfocards(InfocardKey(market_good.Nickname), equip.IdsInfo.Get())
			}

		}
		if equip, ok := e.configs.Equip.ItemsMap[wreck]; ok {
			market_good.Name = e.GetInfocardName(equip.IdsName.Get(), wreck)
			e.exportInfocards(InfocardKey(market_good.Nickname), equip.IdsInfo.Get())
		}

		market_good_key := GetCommodityKey(market_good.Nickname, market_good.ShipClass)
		base.MarketGoodsPerNick[market_good_key] = market_good
	}

	var sb []string
	sb = append(sb, base.Name)
	sb = append(sb, `This is only pseudo base to show availability of lootable content`)
	sb = append(sb, `The content is findable in wrecks or drops from ships at missions`)

	e.Infocards[InfocardKey(base.Nickname)] = sb

	bases = append(bases, base)
	return bases
}

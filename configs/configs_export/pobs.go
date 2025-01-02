package configs_export

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/equip_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/discovery/pob_goods"
	"github.com/darklab8/go-utils/utils/ptr"
)

type ShopItem struct {
	pob_goods.ShopItem
	Nickname string
	Name     string
	Category string
}

type DefenseMode int

func (d DefenseMode) ToStr() string {
	switch d {

	case 1:
		return "SRP Whitelist > Blacklist > IFF Standing, Anyone with good standing"
	case 2:
		return "Whitelist > Nodock, Whitelisted ships only"
	case 3:
		return "Whitelist > Hostile, Whitelisted ships only"
	default:
		return "not recognized"
	}
}

// also known as Player Base Station
type PoB struct {
	Nickname string
	Name     string

	Pos         *string
	Level       *int
	Money       *int
	Health      *int
	DefenseMode *DefenseMode

	SystemNick  *string
	SystemName  *string // SystemHash      *flhash.HashCode `json:"system"`      //: 2745655887,
	FactionNick *string
	FactionName *string // AffiliationHash *flhash.HashCode `json:"affiliation"` //: 2620,

	ShopItems []ShopItem
}

func (e *Exporter) GetPoBs() []PoB {
	var pobs []PoB

	systems_by_hash := make(map[flhash.HashCode]*universe_mapped.System)
	factions_by_hash := make(map[flhash.HashCode]*initialworld.Group)

	for _, system_info := range e.configs.Universe.Systems {
		nickname := system_info.Nickname.Get()
		system_hash := flhash.HashNickname(nickname)
		systems_by_hash[system_hash] = system_info
	}

	for _, group_info := range e.configs.InitialWorld.Groups {
		nickname := group_info.Nickname.Get()
		group_hash := flhash.HashFaction(nickname)
		factions_by_hash[group_hash] = group_info
	}

	goods_by_hash := make(map[flhash.HashCode]*equip_mapped.Item)

	for _, item := range e.configs.Equip.Items {
		nickname := item.Nickname.Get()
		hash := flhash.HashNickname(nickname)
		goods_by_hash[hash] = item
		e.exportInfocards(InfocardKey(nickname), item.IdsInfo.Get())
	}

	for _, pob_info := range e.configs.Discovery.PlayerOwnedBases.Bases {

		var pob PoB = PoB{
			Nickname: pob_info.Nickname,
			Name:     pob_info.Name,
			Pos:      pob_info.Pos,
			Level:    pob_info.Level,
			Money:    pob_info.Money,
			Health:   pob_info.Health,
		}
		if pob_info.DefenseMode != nil {
			pob.DefenseMode = (*DefenseMode)(pob_info.DefenseMode)
		}

		if pob_info.SystemHash != nil {
			if system, ok := systems_by_hash[*pob_info.SystemHash]; ok {
				pob.SystemNick = ptr.Ptr(system.Nickname.Get())
				pob.SystemName = ptr.Ptr(e.GetInfocardName(system.StridName.Get(), system.Nickname.Get()))
			}
		}

		if pob_info.AffiliationHash != nil {
			if faction, ok := factions_by_hash[*pob_info.AffiliationHash]; ok {
				pob.FactionNick = ptr.Ptr(faction.Nickname.Get())
				pob.FactionName = ptr.Ptr(e.GetInfocardName(faction.IdsName.Get(), faction.Nickname.Get()))
			}
		}

		for _, shop_item := range pob_info.ShopItems {
			var good ShopItem = ShopItem{ShopItem: shop_item}

			if item, ok := goods_by_hash[flhash.HashCode(shop_item.Id)]; ok {
				good.Nickname = item.Nickname.Get()
				good.Name = e.GetInfocardName(item.IdsName.Get(), item.Nickname.Get())
				good.Category = item.Category
			}

			pob.ShopItems = append(pob.ShopItems, good)
		}

		pobs = append(pobs, pob)
	}
	return pobs
}

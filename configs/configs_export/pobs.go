package configs_export

import (
	"strconv"
	"strings"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/equip_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"
	"github.com/darklab8/fl-configs/configs/discovery/pob_goods"
	"github.com/darklab8/go-typelog/typelog"
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

		var sb []string
		sb = append(sb, `This is Player Owned Base`)
		sb = append(sb, `infocard placeholder`)

		e.Infocards[InfocardKey(pob.Nickname)] = sb

		pobs = append(pobs, pob)
	}
	return pobs
}

type PobShopItem struct {
	*ShopItem
	PoBName     string
	PobNickname string

	System      *universe_mapped.System
	SystemNick  string
	SystemName  string
	FactionNick string
	FactionName string
	BasePos     *cfgtype.Vector
}

func (e *Exporter) get_pob_buyable() map[string][]*PobShopItem {
	if e.pob_buyable_cache != nil {
		return e.pob_buyable_cache
	}

	e.pob_buyable_cache = make(map[string][]*PobShopItem)

	// TODO refactor copy repeated code may be
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
		for _, shop_item := range pob_info.ShopItems {
			var good *ShopItem = &ShopItem{ShopItem: shop_item}
			if item, ok := goods_by_hash[flhash.HashCode(shop_item.Id)]; ok {
				good.Nickname = item.Nickname.Get()
				good.Name = e.GetInfocardName(item.IdsName.Get(), item.Nickname.Get())
				good.Category = item.Category
			}

			pob_item := &PobShopItem{
				ShopItem:    good,
				PobNickname: pob_info.Nickname,
				PoBName:     pob_info.Name,
			}

			if pob_info.SystemHash != nil {
				if system, ok := systems_by_hash[*pob_info.SystemHash]; ok {
					pob_item.SystemNick = system.Nickname.Get()
					pob_item.SystemName = e.GetInfocardName(system.StridName.Get(), system.Nickname.Get())
					pob_item.System = system
				}
			}
			if pob_info.AffiliationHash != nil {
				if faction, ok := factions_by_hash[*pob_info.AffiliationHash]; ok {
					pob_item.FactionNick = faction.Nickname.Get()
					pob_item.FactionName = e.GetInfocardName(faction.IdsName.Get(), faction.Nickname.Get())
				}
			}

			if pob_info.Pos != nil {
				coords := strings.Split(*pob_info.Pos, ",")
				x, err1 := strconv.ParseFloat(strings.ReplaceAll(coords[0], " ", ""), 64)
				y, err2 := strconv.ParseFloat(strings.ReplaceAll(coords[1], " ", ""), 64)
				z, err3 := strconv.ParseFloat(strings.ReplaceAll(coords[2], " ", ""), 64)
				logus.Log.CheckPanic(err1, "failed parsing x coord", typelog.Any("pos", *pob_info.Pos))
				logus.Log.CheckPanic(err2, "failed parsing y coord", typelog.Any("pos", *pob_info.Pos))
				logus.Log.CheckPanic(err3, "failed parsing z coord", typelog.Any("pos", *pob_info.Pos))

				pob_item.BasePos = &cfgtype.Vector{X: x, Y: y, Z: z}
			}

			e.pob_buyable_cache[good.Nickname] = append(e.pob_buyable_cache[good.Nickname], pob_item)
		}
	}
	return e.pob_buyable_cache
}

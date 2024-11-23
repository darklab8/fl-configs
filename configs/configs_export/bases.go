package configs_export

import (
	"fmt"
	"strings"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped/systems_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/go-utils/utils/ptr"
	"github.com/darklab8/go-utils/utils/utils_types"
)

func VectorToSectorCoord(system *universe_mapped.System, pos cfgtype.Vector) string {
	var scale float64 = 1.0
	if value, ok := system.NavMapScale.GetValue(); ok {
		scale = value
	}

	var fGridSize float64 = 34000.0 / scale // 34000 suspiciously looks like math.MaxInt16
	var gridRefX = int((pos.X+(fGridSize*5))/fGridSize) - 1
	var gridRefZ = int((pos.Z+(fGridSize*5))/fGridSize) - 1
	gridRefX = min(max(gridRefX, 0), 7)
	scXPos := rune('A' + gridRefX)
	gridRefZ = min(max(gridRefZ, 0), 7)
	scZPos := rune('1' + gridRefZ)
	return fmt.Sprintf("%c-%c", scXPos, scZPos)

}

func (e *Exporter) GetBases() []*Base {
	results := make([]*Base, 0, len(e.configs.Universe_config.Bases))

	commodities_per_base := e.getMarketGoods()

	for _, base := range e.configs.Universe_config.Bases {
		var name string = e.GetInfocardName(base.StridName.Get(), base.Nickname.Get())

		var system_name infocard.Infoname
		var Region string
		system, found_system := e.configs.Universe_config.SystemMap[universe_mapped.SystemNickname(base.System.Get())]

		if found_system {

			system_name = infocard.Infoname(e.GetInfocardName(system.Strid_name.Get(), system.Nickname.Get()))

			Region = e.GetRegionName(system)
		}

		var infocard_id int
		var reputation_nickname string
		var pos cfgtype.Vector

		var archetypes []string

		if system, ok := e.configs.Systems.SystemsMap[base.System.Get()]; ok {
			if system_base, ok := system.BasesByBases[base.Nickname.Get()]; ok {
				infocard_id = system_base.IDsInfo.Get()
				reputation_nickname = system_base.RepNickname.Get()
			}

			if system_bases, ok := system.AllBasesByDockWith[base.Nickname.Get()]; ok {
				for _, system_base := range system_bases {
					pos, _ = system_base.Pos.GetValue()
					archetype, _ := system_base.Archetype.GetValue()
					archetypes = append(archetypes, archetype)
				}
			}
		}

		var infocard_ids []int = make([]int, 0)

		infocard_ids = append(infocard_ids, infocard_id)

		if infocard_middle_id, exists := e.configs.InfocardmapINI.InfocardMapTable.Map[infocard_id]; exists {
			infocard_ids = append(infocard_ids, infocard_middle_id)
		}

		var factionName string
		if group, exists := e.configs.InitialWorld.GroupsMap[reputation_nickname]; exists {
			infocard_ids = append(infocard_ids, group.IdsInfo.Get())
			factionName = e.GetInfocardName(group.IdsName.Get(), reputation_nickname)
		}

		var market_goods_per_good_nick map[CommodityKey]MarketGood = make(map[CommodityKey]MarketGood)
		if found_commodities, ok := commodities_per_base[cfgtype.BaseUniNick(base.Nickname.Get())]; ok {
			market_goods_per_good_nick = found_commodities
		}

		var nickname cfgtype.BaseUniNick = cfgtype.BaseUniNick(base.Nickname.Get())

		e.exportInfocards(InfocardKey(nickname), infocard_ids...)

		base := &Base{
			Name:               name,
			Nickname:           nickname,
			NicknameHash:       flhash.HashNickname(nickname.ToStr()),
			FactionName:        factionName,
			System:             string(system_name),
			SystemNickname:     base.System.Get(),
			SystemNicknameHash: flhash.HashNickname(base.System.Get()),
			StridName:          base.StridName.Get(),
			InfocardID:         infocard_id,
			Infocard:           InfocardKey(nickname),
			File:               utils_types.FilePath(base.File.Get()),
			BGCS_base_run_by:   base.BGCS_base_run_by.Get(),
			MarketGoodsPerNick: market_goods_per_good_nick,
			Pos:                pos,
			Archetypes:         archetypes,
			Region:             Region,
		}

		e.Hashes[string(base.Nickname)] = base.NicknameHash
		e.Hashes[base.SystemNickname] = base.SystemNicknameHash

		if found_system {
			base.SectorCoord = VectorToSectorCoord(system, base.Pos)
		}

		results = append(results, base)
	}

	return results
}

func EnhanceBasesWithServerOverrides(bases []*Base, commodities []*Commodity) {
	var base_per_nick map[cfgtype.BaseUniNick]*Base = make(map[cfgtype.BaseUniNick]*Base)
	for _, base := range bases {
		base_per_nick[base.Nickname] = base
	}

	for _, commodity := range commodities {
		for _, base_location := range commodity.Bases {

			// no idea why :) but crashes otherwise
			if base_location.BaseNickname == pob_crafts_nickname {
				continue
			}

			var market_good MarketGood
			commodity_key := GetCommodityKey(commodity.Nickname, commodity.ShipClass)

			if base, ok := base_per_nick[base_location.BaseNickname]; ok {
				if good, ok := base.MarketGoodsPerNick[commodity_key]; ok {
					market_good = good
				}
			}

			market_good.Name = commodity.Name
			market_good.Nickname = commodity.Nickname
			market_good.NicknameHash = commodity.NicknameHash
			market_good.Type = "commodity"

			market_good.LevelRequired = base_location.LevelRequired
			market_good.RepRequired = base_location.RepRequired
			market_good.BaseSells = base_location.BaseSells

			market_good.PriceToBuy = base_location.PriceBaseSellsFor
			market_good.PriceToSell = ptr.Ptr(base_location.PriceBaseBuysFor)
			market_good.Infocard = commodity.Infocard
			market_good.Volume = commodity.Volume
			market_good.IsServerSideOverride = base_location.IsServerSideOverride

			base_per_nick[base_location.BaseNickname].MarketGoodsPerNick[commodity_key] = market_good
		}
	}
}

func FilterToUserfulBases(bases []*Base) []*Base {
	var useful_bases []*Base = make([]*Base, 0, len(bases))
	for _, item := range bases {
		if (item.Name == "Object Unknown" || item.Name == "") && len(item.MarketGoodsPerNick) == 0 {
			continue
		}

		if strings.Contains(item.System, "Bastille") {
			continue
		}

		is_invisible := true
		for _, archetype := range item.Archetypes {
			if archetype != systems_mapped.BaseArchetypeInvisible {
				is_invisible = false
			}
		}
		if is_invisible {
			continue
		}
		item.Reachable = true
		useful_bases = append(useful_bases, item)
	}
	return useful_bases
}

type Base struct {
	Name               string
	Archetypes         []string
	Nickname           cfgtype.BaseUniNick
	NicknameHash       flhash.HashCode
	FactionName        string
	System             string
	SystemNickname     string
	SystemNicknameHash flhash.HashCode
	Region             string
	StridName          int
	InfocardID         int
	Infocard           InfocardKey
	File               utils_types.FilePath
	BGCS_base_run_by   string
	MarketGoodsPerNick map[CommodityKey]MarketGood
	Pos                cfgtype.Vector
	SectorCoord        string

	IsTransportUnreachable bool

	Missions BaseMissions
	BaseAllTradeRoutes
	BaseAllRoutes
	MiningInfo

	Reachable bool
}

type CommodityKey string

func GetCommodityKey(nickname string, ship_class cfgtype.ShipClass) CommodityKey {
	return CommodityKey(fmt.Sprintf("%s_%d", nickname, ship_class))
}

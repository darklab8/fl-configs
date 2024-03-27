package configs_export

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/fl-configs/configs/lower_map"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type NoNameIncluded bool

func (e *Exporter) getBases(is_no_name_included NoNameIncluded) []Base {
	var results []Base = make([]Base, len(e.configs.Universe_config.Bases))

	commodities_per_base := lower_map.NewKeyLoweredMap(lower_map.WithData(e.getMarketGoods()))

	iterator := 0
	for _, base := range e.configs.Universe_config.Bases {
		var name string
		if base_infocard, ok := e.configs.Infocards.Infonames[base.StridName.Get()]; ok {
			name = string(base_infocard)
		}

		if !is_no_name_included && name == "" {
			continue
		}

		var system_name infocard.Infoname
		if system, ok := e.configs.Universe_config.SystemMap.MapGetValue(universe_mapped.SystemNickname(base.System.Get())); ok {

			if infoname, ok := e.configs.Infocards.Infonames[system.Strid_name.Get()]; ok {
				system_name = infoname
			}
		}

		var infocard_id int
		var reputation_nickname string

		if system, ok := e.configs.Systems.SystemsMap.MapGetValue(base.System.Get()); ok {
			for _, system_base := range system.Bases {
				if system_base.IdsName.Get() == base.StridName.Get() {
					infocard_id = system_base.IDsInfo.Get()
					reputation_nickname = system_base.RepNickname.Get()
				}
			}
		}

		var infocardStart []string
		if base_infocard_part1, exists := e.configs.Infocards.Infocards[infocard_id]; exists {
			var err error
			infocardStart, err = base_infocard_part1.XmlToText()
			logus.Log.CheckError(err, "failed to xml infocard")
		}

		var infocardMiddle []string
		if infocard_middle_id, exists := e.configs.InfocardmapINI.InfocardMapTable.Map[infocard_id]; exists {
			if base_infocard_part2, info_exists := e.configs.Infocards.Infocards[infocard_middle_id]; info_exists {
				var err error
				infocardMiddle, err = base_infocard_part2.XmlToText()
				logus.Log.CheckError(err, "failed to xml infocard")
			}
		}

		var infocardEnd []string
		var factionName string
		if group, exists := e.configs.InitialWorld.GroupsMap.MapGetValue(reputation_nickname); exists {
			if infocard_part, info_exists := e.configs.Infocards.Infocards[group.IdsInfo.Get()]; info_exists {
				var err error
				infocardEnd, err = infocard_part.XmlToText()
				logus.Log.CheckError(err, "failed to xml infocard")
			}

			if faction_name, exists := e.configs.Infocards.Infonames[group.IdsName.Get()]; exists {
				factionName = string(faction_name)
			}
		}

		var market_goods []MarketGood
		if found_commodities, ok := commodities_per_base.MapGetValue(base.Nickname.Get()); ok {
			market_goods = found_commodities
		}
		results[iterator] = Base{
			Name:           name,
			Nickname:       base.Nickname.Get(),
			FactionName:    factionName,
			System:         string(system_name),
			SystemNickname: base.System.Get(),
			StridName:      base.StridName.Get(),
			InfocardID:     infocard_id,
			Infocard: Infocard{
				Lines: append(infocardStart, append(infocardMiddle, infocardEnd...)...),
			},
			File:             utils_types.FilePath(base.File.Get()),
			BGCS_base_run_by: base.BGCS_base_run_by.Get(),
			MarketGoods:      market_goods,
		}
		iterator += 1
	}

	results = results[:iterator]

	return results
}

type Infocard struct {
	Lines []string
}

type Base struct {
	Name             string
	Nickname         string
	FactionName      string
	System           string
	SystemNickname   string
	StridName        int
	InfocardID       int
	Infocard         Infocard
	File             utils_types.FilePath
	BGCS_base_run_by string
	MarketGoods      []MarketGood
}

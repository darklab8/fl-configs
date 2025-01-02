package configs_export

import (
	"fmt"
	"strings"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
)

func (e *Exporter) pob_produced() map[string]bool {
	if e.craftable_cached != nil {
		return e.craftable_cached
	}

	e.craftable_cached = make(map[string]bool)

	if e.configs.Discovery != nil {
		for _, recipe := range e.configs.Discovery.BaseRecipeItems.Recipes {
			for _, produced := range recipe.ProcucedItem {
				e.craftable_cached[produced.Get()] = true
			}
		}
	}

	if e.configs.FLSR != nil {
		for _, recipe := range e.configs.FLSR.FLSRRecipes.Products {
			e.craftable_cached[recipe.Product.Get()] = true
		}
	}

	return e.craftable_cached
}

const (
	pob_crafts_nickname = "crafts"
)

func (e *Exporter) CraftableBaseName() string {
	if e.configs.Discovery != nil {
		return "PoB crafts"
	}
	if e.configs.FLSR != nil {
		return "Craftable"
	}

	return "NoCrafts"
}

func (e *Exporter) EnhanceBasesWithPobCrafts(bases []*Base) []*Base {
	pob_produced := e.pob_produced()

	base := &Base{
		Name:               e.CraftableBaseName(),
		MarketGoodsPerNick: make(map[CommodityKey]MarketGood),
		Nickname:           cfgtype.BaseUniNick(pob_crafts_nickname),
		Infocard:           InfocardKey(pob_crafts_nickname),
		SystemNickname:     "neverwhere",
		System:             "Neverwhere",
		Region:             "Neverwhere",
		FactionName:        "Player Crafts",
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

		if good, found_good := e.configs.Goods.GoodsMap[market_good.Nickname]; found_good {
			category := good.Category.Get()
			market_good.Type = fmt.Sprintf("%s craft", category)
			if equip, ok := e.configs.Equip.ItemsMap[market_good.Nickname]; ok {
				market_good.Type = fmt.Sprintf("%s craft", equip.Category)
				e.exportInfocards(InfocardKey(market_good.Nickname), equip.IdsInfo.Get())
			}

		}
		var ship_nickname string
		if good_ship, ok := e.configs.Goods.ShipsMap[produced]; ok {
			hull_name := good_ship.Hull.Get()
			if good_shiphull, ok := e.configs.Goods.ShipHullsMap[hull_name]; ok {
				ship_nick := good_shiphull.Ship.Get()
				ship_nickname = ship_nick
				if equipment, ok := e.configs.Shiparch.ShipsMap[ship_nick]; ok {
					market_good.Name = e.GetInfocardName(equipment.IdsName.Get(), market_good.Nickname)
					e.exportInfocards(InfocardKey(market_good.Nickname), equipment.IdsInfo1.Get(), equipment.IdsInfo.Get())
				}
			}
		} else {
			if equip, ok := e.configs.Equip.ItemsMap[produced]; ok {
				market_good.Name = e.GetInfocardName(equip.IdsName.Get(), produced)
				e.exportInfocards(InfocardKey(market_good.Nickname), equip.IdsInfo.Get())
			}
		}

		market_good_key := GetCommodityKey(market_good.Nickname, market_good.ShipClass)
		base.MarketGoodsPerNick[market_good_key] = market_good

		var infocard_addition []string
		if e.configs.Discovery != nil {
			if recipes, ok := e.configs.Discovery.BaseRecipeItems.RecipePerProduced[market_good.Nickname]; ok {
				infocard_addition = append(infocard_addition, `CRAFTING RECIPES:`)
				for _, recipe := range recipes {
					sector := recipe.Model.RenderModel()
					infocard_addition = append(infocard_addition, string(sector.OriginalType))
					for _, param := range sector.Params {
						infocard_addition = append(infocard_addition, string(param.ToString(inireader.WithComments(false))))
					}
					infocard_addition = append(infocard_addition, "")
				}
			}
		}
		if e.configs.FLSR != nil {
			if e.configs.FLSR.FLSRRecipes != nil {
				if recipes, ok := e.configs.FLSR.FLSRRecipes.ProductsByNick[market_good.Nickname]; ok {
					infocard_addition = append(infocard_addition, `CRAFTING RECIPES:`)
					for _, recipe := range recipes {
						sector := recipe.Model.RenderModel()
						infocard_addition = append(infocard_addition, string(sector.OriginalType))
						for _, param := range sector.Params {
							infocard_addition = append(infocard_addition, string(param.ToString(inireader.WithComments(false))))
						}
						infocard_addition = append(infocard_addition, "")
					}
				}
			}
		}

		var info Infocard
		if value, ok := e.Infocards[market_good.Infocard]; ok {
			info = value
		}

		add_line_about_recipes := func(info Infocard) Infocard {
			add_line := func(index int, line string) {
				info = append(info[:index+1], info[index:]...)
				info[index] = line
			}
			strip_line := func(line string) string {
				return strings.ReplaceAll(strings.ReplaceAll(line, " ", ""), "\u00a0", "")
			}
			if len(infocard_addition) > 0 {
				line_position := 1
				add_line(line_position, `<b>Item has crafting recipes below</b>`)
				if strip_line(info[0]) != "" {
					add_line(1, "")
					line_position += 1
				}
				if strip_line(info[line_position+1]) != "" {
					add_line(line_position+1, "")
				}
			}
			return info
		}
		info = add_line_about_recipes(info)

		e.Infocards[market_good.Infocard] = append(info, infocard_addition...)

		if ship_nickname != "" {
			var info Infocard
			if value, ok := e.Infocards[InfocardKey(ship_nickname)]; ok {
				info = value
			}
			info = add_line_about_recipes(info)
			e.Infocards[InfocardKey(ship_nickname)] = append(info, infocard_addition...)
		}
	}

	var sb []string
	sb = append(sb, base.Name)
	sb = append(sb, `This is only pseudo base to show availability of player crafts`)
	sb = append(sb, ``)
	sb = append(sb, `At the bottom of each item infocard it shows CRAFTING RECIPES`)

	e.Infocards[InfocardKey(base.Nickname)] = sb

	bases = append(bases, base)
	return bases
}

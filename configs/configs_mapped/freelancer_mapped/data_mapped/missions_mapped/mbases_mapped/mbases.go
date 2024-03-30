package mbases_mapped

import (
	"fmt"
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
)

const (
	FILENAME = "mbases.ini"
)

type Mroom struct {
	semantic.Model
	Nickname         *semantic.String
	CharacterDensity *semantic.Int
	Bartrender       *semantic.String
}

type BaseFaction struct {
	semantic.Model

	Faction *semantic.String
	Weight  *semantic.Int
	Npcs    []*semantic.String
}

type Bribe struct {
	semantic.Model
	Faction *semantic.String
}
type Rumor struct {
	semantic.Model
}
type Mission struct {
	semantic.Model
}
type Know struct {
	semantic.Model
}

type NPC struct {
	semantic.Model

	Nickname    *semantic.String
	Room        *semantic.String
	Bribes      []*Bribe
	Rumors      []*Rumor
	Missions    []*Mission
	Knows       []*Know
	Affiliation *semantic.String
}

type Base struct {
	semantic.Model

	Nickname     *semantic.String
	LocalFaction *semantic.String
	Diff         *semantic.Int

	BaseFactions    []*BaseFaction
	BaseFactionsMap *lower_map.KeyLoweredMap[string, *BaseFaction]
	NPCs            []*NPC
	Bar             *Mroom
}

type Config struct {
	semantic.ConfigModel

	Bases   []*Base
	BaseMap *lower_map.KeyLoweredMap[string, *Base]
}

func Read(input_file *iniload.IniLoader) *Config {
	frelconfig := &Config{
		Bases:   make([]*Base, 0, 100),
		BaseMap: lower_map.NewKeyLoweredMap[string, *Base](),
	}

	for i := 0; i < len(input_file.Sections); i++ {

		if input_file.Sections[i].Type == "[MBase]" {

			mbase_section := input_file.Sections[i]
			base := &Base{
				BaseFactionsMap: lower_map.NewKeyLoweredMap[string, *BaseFaction](),
			}
			base.Map(mbase_section)
			base.Nickname = semantic.NewString(mbase_section, "nickname")
			base.LocalFaction = semantic.NewString(mbase_section, "local_faction")
			base.Diff = semantic.NewInt(mbase_section, "diff")
			frelconfig.Bases = append(frelconfig.Bases, base)
			frelconfig.BaseMap.MapSet(base.Nickname.Get(), base)

			for j := i + 1; j < len(input_file.Sections) && input_file.Sections[j].Type != "[MBase]"; j++ {
				section := input_file.Sections[j]

				switch section.Type {
				case "[BaseFaction]":
					faction := &BaseFaction{
						Faction: semantic.NewString(section, "faction"),
						Weight:  semantic.NewInt(section, "weight"),
					}
					faction.Map(section)

					for index, _ := range section.ParamMap["npc"] {
						faction.Npcs = append(faction.Npcs,
							semantic.NewString(mbase_section, "weight", semantic.OptsS(semantic.Index(index))))
					}
					base.BaseFactions = append(base.BaseFactions, faction)
					base.BaseFactionsMap.MapSet(faction.Faction.Get(), faction)
				case "[MRoom]":
					mroom := &Mroom{
						Nickname:         semantic.NewString(section, "nickname"),
						CharacterDensity: semantic.NewInt(section, "character_density"),
						Bartrender:       semantic.NewString(section, "fixture", semantic.OptsS(semantic.Order(0), semantic.Optional())),
					}
					mroom.Map(section)
					if strings.ToLower(mroom.Nickname.Get()) == "bar" {
						base.Bar = mroom
					}
				case "[GF_NPC]":
					npc := &NPC{
						Nickname:    semantic.NewString(section, "nickname"),
						Room:        semantic.NewString(section, "room", semantic.OptsS(semantic.Optional())),
						Affiliation: semantic.NewString(section, "affiliation"),
					}
					npc.Map(section)

					for index, _ := range section.ParamMap["bribe"] {
						bribe := &Bribe{
							Faction: semantic.NewString(section, "bribe", semantic.OptsS(semantic.Index(index))),
						}
						bribe.Map(section)
						npc.Bribes = append(npc.Bribes, bribe)
					}
					for range section.ParamMap["rumor"] {
						rumor := &Rumor{}
						rumor.Map(section)
						npc.Rumors = append(npc.Rumors, rumor)
					}
					for range section.ParamMap["misc"] {
						misn := &Mission{}
						misn.Map(section)
						npc.Missions = append(npc.Missions, misn)
					}
					for range section.ParamMap["know"] {
						know := &Know{}
						know.Map(section)
						npc.Knows = append(npc.Knows, know)
					}

					base.NPCs = append(base.NPCs, npc)
				}

			}
		}

	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

type BaseChance struct {
	Base   string
	Chance float64
}

func FactionRephacks(config *Config) map[string]map[string]float64 {
	// for faction, chance at certain base
	var faction_rephacks map[string]map[string]float64 = make(map[string]map[string]float64)

	for _, base := range config.Bases {

		// per faction chance at base
		fmt.Println("base=", base.Nickname.Get())
		var base_bribe_chances map[string]float64 = make(map[string]float64)
		var faction_members map[string]int = make(map[string]int)
		for _, npc := range base.NPCs {
			faction_members[strings.ToLower(npc.Affiliation.Get())] += 1

		}
		for _, npc := range base.NPCs {
			if base.Bar == nil {
				continue
			}
			npc_nickname := npc.Nickname.Get()
			bartrender := base.Bar.Bartrender.Get()
			if npc_nickname == bartrender {
				for _, bribe := range npc.Bribes {
					chance_increase := 1 / float64(len(npc.Bribes)+len(npc.Rumors)+len(npc.Missions)+len(npc.Knows))
					base_bribe_chances[strings.ToLower(bribe.Faction.Get())] += chance_increase
				}
			} else {
				for _, bribe := range npc.Bribes {
					var weight float64 = 0
					if faction, ok := base.BaseFactionsMap.MapGetValue(npc.Affiliation.Get()); ok {
						weight = float64(faction.Weight.Get())

						if value, ok := faction_members[strings.ToLower(npc.Affiliation.Get())]; ok {
							if value != 0 {
								weight = weight / float64(value)
							}
						}
					}

					chance_increase := float64(weight/100) * 1 / float64(len(npc.Bribes)+len(npc.Rumors)+len(npc.Missions)+len(npc.Knows))
					base_bribe_chances[strings.ToLower(bribe.Faction.Get())] += chance_increase
				}
			}
		}

		for faction, chance := range base_bribe_chances {
			_, ok := faction_rephacks[strings.ToLower(faction)]
			if !ok {
				faction_rephacks[strings.ToLower(faction)] = make(map[string]float64)
			}
			faction_rephacks[strings.ToLower(faction)][strings.ToLower(base.Nickname.Get())] += chance

			if faction_rephacks[strings.ToLower(faction)][strings.ToLower(base.Nickname.Get())] > 1.0 {
				faction_rephacks[strings.ToLower(faction)][strings.ToLower(base.Nickname.Get())] = 1.0
			}
		}

	}
	return faction_rephacks
}

package mbases_mapped

import (
	"fmt"
	"sort"
	"strings"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/stretchr/testify/assert"
)

func TestGetRepHacks(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_filepath.Join(test_directory, FILENAME))

	config := Read(iniload.NewLoader(fileref).Scan())
	assert.Greater(t, len(config.Bases), 0, "expected finding some elements")

	// for faction, chance at certain base
	var faction_rephacks map[string]map[string]float64 = make(map[string]map[string]float64)

	for index, base := range config.Bases {

		// per faction chance at base
		fmt.Println("base=", base.Nickname.Get())
		var base_bribe_chances map[string]float64 = make(map[string]float64)
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
					}

					chance_increase := float64(weight/100) * 1 / float64(len(npc.Bribes)+len(npc.Rumors)+len(npc.Missions)+len(npc.Knows))
					base_bribe_chances[strings.ToLower(bribe.Faction.Get())] += chance_increase
				}
			}
		}

		for faction, chance := range base_bribe_chances {
			_, ok := faction_rephacks[faction]
			if !ok {
				faction_rephacks[faction] = make(map[string]float64)
			}
			faction_rephacks[faction][base.Nickname.Get()] += chance
		}

		if index == 0 {
			fmt.Println("chances")
			for faction, chance := range base_bribe_chances {
				fmt.Println(faction, " = ", chance*100)
			}
		}
	}

	fmt.Println("printing for br_p_grp")
	chances := make([]BaseChance, 0, len(faction_rephacks["br_p_grp"]))

	for base, chance := range faction_rephacks["br_p_grp"] {
		chances = append(chances, BaseChance{
			base:   base,
			chance: chance,
		})
	}
	sort.Slice(chances, func(i, j int) bool {
		return chances[i].chance > chances[j].chance
	})

	for _, chance := range chances {
		fmt.Println(chance.base, " = ", 100*chance.chance)
	}
}

type BaseChance struct {
	base   string
	chance float64
}

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

	// configs := configs_mapped.TestFixtureConfigs()
	// exporter := configs_export.NewExporter(configs)
	// bases := exporter.GetBases(configs_export.NoNameIncluded(true))

	faction_rephacks := FactionRephacks(config)

	fmt.Println("printing for br_p_grp")
	chances := make([]BaseChance, 0, len(faction_rephacks[strings.ToLower("br_p_grp")]))

	for base, chance := range faction_rephacks[strings.ToLower("br_p_grp")] {
		chances = append(chances, BaseChance{
			base:   base,
			chance: chance,
		})
	}
	sort.Slice(chances, func(i, j int) bool {
		return chances[i].chance > chances[j].chance
	})

	for _, chance := range chances {
		var name string
		// for _, base := range bases {
		// 	if strings.ToLower(chance.base) == strings.ToLower(base.Nickname) {
		// 		name = base.Name
		// 	}
		// }

		fmt.Println(chance.base, " = ", 100*chance.chance, " ", name)
	}
}

type BaseChance struct {
	base   string
	chance float64
}

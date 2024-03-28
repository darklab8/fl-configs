package equipment_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_filepath.Join(test_directory, FILENAME))

	config := Read([]*file.File{fileref})

	assert.Greater(t, len(config.Commodities), 0, "expected finding commodities")
	assert.Greater(t, len(config.ShipHulls), 0)
	assert.Greater(t, len(config.Ships), 0)

	for _, commodity := range config.Commodities {
		commodity.Price.Get()
	}
}

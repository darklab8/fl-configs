package const_mapped

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref1 := file.NewFile(utils_filepath.Join(test_directory, "constants.ini"))
	config := Read(iniload.NewLoader(fileref1).Scan())
	fmt.Println(config.EngineEquipConsts.CRUISING_SPEED.Get())
}

func TestBiniReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	filepath := utils_filepath.Join(test_directory, "constants.vanilla.ini")

	fileref1 := file.NewFile(filepath)
	config := Read(iniload.NewLoader(fileref1).Scan())
	fmt.Println(config.ShieldEquipConsts.HULL_DAMAGE_FACTOR.Get())
}

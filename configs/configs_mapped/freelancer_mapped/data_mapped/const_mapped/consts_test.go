package const_mapped

import (
	"fmt"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/iniload"
	"github.com/darklab8/fl-configs/configs/tests"
)

func TestReader(t *testing.T) {
	fileref1 := tests.FixtureFileFind().GetFile("constants.ini")
	config := Read(iniload.NewLoader(fileref1).Scan())
	fmt.Println(config.EngineEquipConsts.CRUISING_SPEED.Get())
}

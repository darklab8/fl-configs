package infocard_mapped

import (
	"path/filepath"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"github.com/stretchr/testify/assert"
)

// Not used any longer?
func TestReader(t *testing.T) {
	one_file_filesystem := filefind.FindConfigs(utils.GetCurrrentTestFolder())

	config := infocard.NewConfig()
	test_directory, _ := filepath.Abs(utils.GetCurrentFolder().ToString())
	test_directory = filepath.Join(filepath.Dir(test_directory), "exe_mapped", "testdata")
	filesystem := filefind.FindConfigs(utils_types.FilePath(test_directory))

	freelancer_ini := exe_mapped.FixtureFLINIConfig()
	config, _ = Read(filesystem, freelancer_ini, one_file_filesystem.GetFile("temp.disco.infocards.txt"))

	assert.Greater(t, len(config.Infocards), 0)
}

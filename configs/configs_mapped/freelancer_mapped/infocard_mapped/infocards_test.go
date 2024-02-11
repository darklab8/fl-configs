package infocard_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"

	"github.com/darklab8/go-utils/goutils/utils"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	config := &Config{}
	filesystem := filefind.FindConfigs(test_directory)

	freelancer_ini := exe_mapped.FixtureFLINIConfig()
	config.Read(filesystem, freelancer_ini, filesystem.GetFile(FILENAME))

	assert.Greater(t, len(config.Records), 0)
}

package infocard_mapped

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_filepath"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	test_directory := utils.GetCurrrentTestFolder()
	fileref := file.NewFile(utils_filepath.Join(test_directory, FILENAME))
	config := Config{}
	config.Read(fileref)

	assert.Greater(t, len(config.Records), 0)
}

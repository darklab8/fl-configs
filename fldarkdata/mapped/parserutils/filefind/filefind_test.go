package filefind

import (
	"testing"

	"github.com/darklab8/darklab_goutils/goutils/utils"

	"github.com/stretchr/testify/assert"
)

func TestDiscoverFiles(t *testing.T) {
	// Write some data example in order to remove integration flag
	test_directory := utils.GetCurrrentTestFolder()
	filesystem := FindConfigs(test_directory)

	assert.Equal(t, 2, len(filesystem.Files), "expected 2 files, fount smth else")
}

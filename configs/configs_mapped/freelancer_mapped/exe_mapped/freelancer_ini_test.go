package exe_mapped

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	config := FixtureFLINIConfig()
	assert.Greater(t, len(config.Resources.Dll), 0)
}

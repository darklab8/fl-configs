package inireader

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/tests"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	fileref := tests.FixtureFileFind().GetFile("market_ships.ini")
	config := Read(fileref)

	assert.Greater(t, len(config.Sections), 0, "market ships sections were not scanned")
}

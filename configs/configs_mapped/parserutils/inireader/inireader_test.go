package inireader

import (
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/tests"
	"github.com/darklab8/go-utils/utils"
	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	fileref := tests.FixtureFileFind().GetFile("market_ships.ini")
	config := Read(fileref)

	assert.Greater(t, len(config.Sections), 0, "market ships sections were not scanned")
}

func TestReaderWithBOMFails(t *testing.T) {

	defer func() {
		InitRegexExpression(&regexSection, regexSectionRegExp)
	}()
	InitRegexExpression(&regexSection, `^\[.*\]`)

	fs := filefind.FindConfigs(utils.GetCurrrentTestFolder())
	fileref := fs.GetFile("li05_with_bom.ini")

	var crashed bool = false
	func() {
		defer func() {
			if r := recover(); r != nil {
				crashed = true
			}
		}()
		Read(fileref)
	}()

	assert.True(t, crashed, "with BOM we Crash.")
}

func TestReaderWithBOMPasses(t *testing.T) {

	fs := filefind.FindConfigs(utils.GetCurrrentTestFolder())
	fileref := fs.GetFile("li05_with_bom.ini")
	config := Read(fileref)

	assert.Greater(t, len(config.Sections), 0, "market ships sections were not scanned")
}

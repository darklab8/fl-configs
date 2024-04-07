package configs_export

import (
	"fmt"
	"strings"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/stretchr/testify/assert"
)

func TestGetShips(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	exporter := NewExporter(configs)

	items := exporter.GetShips()
	assert.Greater(t, len(items), 0)

	filtered := FilterToUsefulShips(items)
	assert.Greater(t, len(filtered), 0)

	for _, item := range items {
		if strings.Contains(item.Nickname, "loki") {
			fmt.Println()
		}
	}
}

package configs_export

import (
	"fmt"
	"sync"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_export/trades"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
)

func TestGetTrades(t *testing.T) {
	configs := configs_mapped.TestFixtureConfigs()
	e := NewExporter(configs)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		e.transport = NewGraphResults(e.configs, trades.WithFreighterPaths(false))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		e.freighter = NewGraphResults(e.configs, trades.WithFreighterPaths(true))
		wg.Done()
	}()
	e.Bases = e.GetBases()
	e.Commodities = e.GetCommodities()
	wg.Wait()
	e.Bases, e.Commodities = e.TradePaths(e.Bases, e.Commodities)

	fmt.Println()
}

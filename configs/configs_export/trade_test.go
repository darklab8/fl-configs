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
		e.transport = NewGraphResults(e, trades.AvgTransportCruiseSpeed, trades.WithFreighterPaths(false))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		e.freighter = NewGraphResults(e, trades.AvgFreighterCruiseSpeed, trades.WithFreighterPaths(true))
		wg.Done()
	}()
	e.Bases = e.GetBases()
	e.Commodities = e.GetCommodities()
	wg.Wait()
	e.Bases, e.Commodities = e.TradePaths(e.Bases, e.Commodities)

	for _, base := range e.Bases {
		for _, trade_route := range base.TradeRoutes {
			trade_route.Transport.GetPaths()
		}
	}

	fmt.Println()
}

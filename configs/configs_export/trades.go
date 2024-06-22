package configs_export

import (
	"github.com/darklab8/fl-configs/configs/configs_export/trades"
)

type Trades struct {
	TradeRoutes        []*ComboTradeRoute
	BestTransportRoute *TradeRoute
	BestFreighterRoute *TradeRoute
}

type TradeRoute struct {
	g           *GraphResults
	Commodity   *Commodity
	BuyingGood  *GoodAtBase
	SellingGood *GoodAtBase
}

type ComboTradeRoute struct {
	Transport *TradeRoute
	Freighter *TradeRoute
}

func NewTradeRoute(g *GraphResults, buying_good *GoodAtBase, selling_good *GoodAtBase, commodity *Commodity) *TradeRoute {
	return &TradeRoute{
		g:           g,
		BuyingGood:  buying_good,
		SellingGood: selling_good,
		Commodity:   commodity,
	}
}

func (t *TradeRoute) GetProffitPerV() float64 {
	return float64(t.SellingGood.PriceBaseBuysFor-t.BuyingGood.PriceBaseSellsFor) / float64(t.Commodity.Volume)
}

func (t *TradeRoute) GetPaths() []string {
	return trades.GetPath(t.g.graph, t.g.parents, t.BuyingGood.BaseNickname, t.SellingGood.BaseNickname)
}

func (t *TradeRoute) GetDist() int {
	return trades.GetDist(t.g.graph, t.g.dists, t.BuyingGood.BaseNickname, t.SellingGood.BaseNickname)
}

func (t *TradeRoute) GetProffitPerTime() float64 {
	return t.GetProffitPerV() / t.GetTime()
}

func (t *TradeRoute) GetTime() float64 {
	return float64(t.GetDist()) / float64(trades.AvgCruiseSpeed)
}

func (e *Exporter) TradePaths(
	bases []*Base,
	commodities []*Commodity,
) ([]*Base, []*Commodity) {

	var commodity_by_nick map[string]*Commodity = make(map[string]*Commodity)
	var commodity_by_good_and_base map[string]map[string]*GoodAtBase = make(map[string]map[string]*GoodAtBase)
	for _, commodity := range commodities {
		commodity_by_nick[commodity.Nickname] = commodity
		if _, ok := commodity_by_good_and_base[commodity.Nickname]; !ok {
			commodity_by_good_and_base[commodity.Nickname] = make(map[string]*GoodAtBase)
		}
		for _, good_at_base := range commodity.Bases {
			commodity_by_good_and_base[commodity.Nickname][good_at_base.BaseNickname] = good_at_base
		}
	}

	for _, base := range bases {
		for _, good := range base.MarketGoods {
			if good.Type != "commodity" {
				continue
			}

			if !good.BaseSells {
				continue
			}

			commodity := commodity_by_nick[good.Nickname]
			buying_good := commodity_by_good_and_base[good.Nickname][base.Nickname]

			if buying_good == nil {
				continue
			}

			for _, selling_good_at_base := range commodity.Bases {
				trade_route := &ComboTradeRoute{
					Transport: NewTradeRoute(e.transport, buying_good, selling_good_at_base, commodity),
					Freighter: NewTradeRoute(e.freighter, buying_good, selling_good_at_base, commodity),
				}

				if trade_route.Transport.GetProffitPerV() <= 0 {
					continue
				}

				// path := trade_route.Transport.GetPaths()
				// _ = path
				// count tradelanes. count jumpholes. Add jumpholes
				// trades.GetPath(graph, parents, "li01_01_base", "br01_01_base")

				base.TradeRoutes = append(base.TradeRoutes, trade_route)
				commodity.TradeRoutes = append(commodity.TradeRoutes, trade_route)
			}
		}
		// bases[base_i]
	}

	for _, commodity := range commodities {
		for _, trade_route := range commodity.TradeRoutes {
			if commodity.BestTransportRoute == nil {
				commodity.BestTransportRoute = trade_route.Transport
			} else if trade_route.Transport.GetProffitPerTime() > commodity.BestTransportRoute.GetProffitPerTime() {
				commodity.BestTransportRoute = trade_route.Transport
			}

			if commodity.BestFreighterRoute == nil {
				commodity.BestFreighterRoute = trade_route.Freighter
			} else if trade_route.Freighter.GetProffitPerTime() > commodity.BestFreighterRoute.GetProffitPerTime() {
				commodity.BestFreighterRoute = trade_route.Freighter
			}
		}
	}

	for _, base := range bases {
		for _, trade_route := range base.TradeRoutes {
			if base.BestTransportRoute == nil {
				base.BestTransportRoute = trade_route.Transport
			} else if trade_route.Transport.GetProffitPerTime() > base.BestTransportRoute.GetProffitPerTime() {
				base.BestTransportRoute = trade_route.Transport
			}

			if base.BestFreighterRoute == nil {
				base.BestFreighterRoute = trade_route.Freighter
			} else if trade_route.Freighter.GetProffitPerTime() > base.BestFreighterRoute.GetProffitPerTime() {
				base.BestFreighterRoute = trade_route.Freighter
			}
		}
	}

	return bases, commodities
}

package configs_export

type MarketGood struct {
	GoodNickname  string
	PriceModifier float64
	//Type Commidity Engine CD and etc
	// Base Sells
	// Level req
	// Reputation Req
	// Price Modifier
	// Nickname

	// Price 364 with 2.0278 for criminals at Alcatran Deplot, li02_06_base
	// market_commidities, price modifier 2.0277777777777777
	// goods.ini price = 180
	// final price, floored down Price * Price modifier
}

func (e *Exporter) GetMarketGoods() []MarketGood {
	var Goods []MarketGood = make([]MarketGood, 0, 20)

	for _, base_good := range e.configs.MarketCommidities.BaseGoods {
		base_nickname := base_good.Base.Get()
		_ = base_nickname

		for _, market_good := range base_good.MarketGoods {
			Goods = append(Goods, MarketGood{
				GoodNickname:  market_good.Nickname.Get(),
				PriceModifier: market_good.PriceModifier.Get(),
			})
		}

	}
	return Goods
}

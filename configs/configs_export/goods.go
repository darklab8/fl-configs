package configs_export

type GoodType string

const (
	TypeCommodity GoodType = "commodity"
)

type MarketGood struct {
	GoodNickname string
	Type         GoodType

	LevelRequired int
	RepRequired   float64

	IsBuyOnly     bool
	PriceModifier float64
	PriceBase     float64
	Price         float64
}

func (e *Exporter) GetMarketGoods() []MarketGood {
	var Goods []MarketGood = make([]MarketGood, 0, 20)

	for _, base_good := range e.configs.MarketCommidities.BaseGoods {
		base_nickname := base_good.Base.Get()
		_ = base_nickname

		for _, market_good := range base_good.MarketGoods {
			Goods = append(Goods, MarketGood{
				GoodNickname:  market_good.Nickname.Get(),
				Type:          TypeCommodity,
				LevelRequired: market_good.LevelRequired.Get(),
				RepRequired:   market_good.RepRequired.Get(),
				IsBuyOnly:     market_good.IsBuyOnly.Get(),
				PriceModifier: market_good.PriceModifier.Get(),

				// TODO
				// PriceBase:
				// Price:
			})
		}

	}
	return Goods
}

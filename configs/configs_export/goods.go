package configs_export

type GoodType string

const (
	TypeCommodity GoodType = "commodity"
)

type MarketGood struct {
	Name     string
	Nickname string
	Type     GoodType

	LevelRequired int
	RepRequired   float64

	IsBuyOnly     bool
	PriceModifier float64
	PriceBase     float64
	Price         float64
}

func (e *Exporter) GetMarketGoods() map[string][]MarketGood {
	var GoodsPerBase map[string][]MarketGood = make(map[string][]MarketGood)

	for _, base_good := range e.configs.MarketCommidities.BaseGoods {
		base_nickname := base_good.Base.Get()

		var MarketGoods []MarketGood = make([]MarketGood, 0, 20)
		for _, market_good := range base_good.MarketGoods {

			// var name string
			// if base_infocard, ok := e.configs.Infocards.Infonames[base.StridName.Get()]; ok {
			// 	name = string(base_infocard)
			// }

			MarketGoods = append(MarketGoods, MarketGood{
				Nickname:      market_good.Nickname.Get(),
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

		GoodsPerBase[base_nickname] = MarketGoods

	}
	return GoodsPerBase
}

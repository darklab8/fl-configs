package configs_export

type Ship struct {
	Name      string
	Class     int
	Type      string
	Price     int
	Armor     int
	HoldSize  int
	Nanobots  int
	Batteries int
	Mass      float64

	PowerCapacity     int
	PowerRechargeRate int
	CruiseSpeed       int
	ImpulseSpeed      float64
	ReverseFraction   float64
	ThrustCapacity    int
	ThrustRecharge    int

	// MaxAngularSpeedDegS           float64
	// AngularDistanceFrom0ToHalfSec float64
	// TimeTo90MaxAngularSpeed       float64

	NudgeForce  float64
	StrafeForce float64
	Nickname    string
	NameID      int
	InfoID      int

	Bases []GoodAtBase
}

func (e *Exporter) GetShips() []Ship {
	var ships []Ship

	for _, ship_info := range e.configs.Shiparch.Ships {
		ship := Ship{
			Nickname: ship_info.Nickname.Get(),
		}

		ship.Class, _ = ship_info.ShipClass.GetValue()
		ship.Type = ship_info.Type.Get()
		ship.NameID = ship_info.IdsName.Get()
		ship.InfoID, _ = ship_info.IdsInfo.GetValue()

		if bots, ok := ship_info.Nanobots.GetValue(); ok {
			ship.Nanobots = bots
		} else {
			continue
		}
		ship.Batteries = ship_info.Batteries.Get()
		ship.Mass = ship_info.Mass.Get()
		ship.NudgeForce = ship_info.NudgeForce.Get()
		ship.StrafeForce, _ = ship_info.StrafeForce.GetValue()

		if name, ok := e.configs.Infocards.Infonames[ship.NameID]; ok {
			ship.Name = string(name)
		}

		if ship_hull_good, ok := e.configs.Goods.ShipHullsMapByShip[ship.Nickname]; ok {
			ship.Price = ship_hull_good.Price.Get()

			ship_hull_nickname := ship_hull_good.Nickname.Get()
			if ship_package_good, ok := e.configs.Goods.ShipsMapByHull[ship_hull_nickname]; ok {
				ship.Bases = e.GetAtBasesSold(GetAtBasesInput{
					Nickname:       ship_package_good.Nickname.Get(),
					Price:          ship.Price,
					PricePerVolume: -1,
				})

				for _, addon := range ship_package_good.Addons {
					// can be Power or Engine or Smth else
					// addon = dsy_hessian_engine, HpEngine01, 1
					// addon = dsy_loki_core, internal, 1
					// addon = ge_s_scanner_01, internal, 1
					addon_nickname := addon.ItemNickname.Get()

					if power, ok := e.configs.Equip.PowersMap[addon_nickname]; ok {
						ship.PowerCapacity = power.Capacity.Get()
						ship.PowerRechargeRate = power.ChargeRate.Get()

						ship.ThrustCapacity = power.ThrustCapacity.Get()
						ship.ThrustRecharge = power.ThrustRecharge.Get()
					}
					if engine, ok := e.configs.Equip.EnginesMap[addon_nickname]; ok {
						ship.CruiseSpeed, _ = engine.CruiseSpeed.GetValue()

						engine_linear_drag, _ := engine.LinearDrag.GetValue()
						ship_linear_drag, _ := ship_info.LinearDrag.GetValue()
						engine_max_force, _ := engine.MaxForce.GetValue()
						ship.ImpulseSpeed = float64(engine_max_force) / (float64(engine_linear_drag) + float64(ship_linear_drag))

						ship.ReverseFraction = engine.ReverseFraction.Get()
					}
				}
			}

		}

		ship.HoldSize = ship_info.HoldSize.Get()
		ship.Armor = ship_info.HitPts.Get()

		if len(ship.Bases) == 0 {
			continue
		}

		ships = append(ships, ship)
	}

	return ships
}

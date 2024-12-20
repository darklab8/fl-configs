package configs_export

import (
	"math"
	"sort"
	"strings"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"
	"github.com/darklab8/go-typelog/typelog"
)

func (g Ship) GetNickname() string                 { return g.Nickname }
func (g Ship) GetTechCompat() *DiscoveryTechCompat { return g.DiscoveryTechCompat }

type Ship struct {
	Nickname     string
	NicknameHash flhash.HashCode

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

	MaxAngularSpeedDegS           float64
	AngularDistanceFrom0ToHalfSec float64
	TimeTo90MaxAngularSpeed       float64

	NudgeForce  float64
	StrafeForce float64
	NameID      int
	InfoID      int

	Bases            map[cfgtype.BaseUniNick]*GoodAtBase
	Slots            []EquipmentSlot
	BiggestHardpoint []string

	*DiscoveryTechCompat

	DiscoShip *DiscoShip
}

type DiscoShip struct {
	ArmorMult float64
}

func (e *Exporter) GetShips(ids []Tractor, TractorsByID map[cfgtype.TractorID]Tractor) []Ship {
	var ships []Ship

	for _, ship_info := range e.configs.Shiparch.Ships {
		ship := Ship{
			Nickname: ship_info.Nickname.Get(),
			Bases:    make(map[cfgtype.BaseUniNick]*GoodAtBase),
		}
		ship.NicknameHash = flhash.HashNickname(ship.Nickname)
		e.Hashes[ship.Nickname] = ship.NicknameHash

		// defer func() {
		// 	if r := recover(); r != nil {
		// 		fmt.Println("Recovered in f", r)
		// 		fmt.Println("ship.Nickname", ship.Nickname)
		// 		panic(r)
		// 	}
		// }()

		ship.Class, _ = ship_info.ShipClass.GetValue()
		if _, ok := ship_info.Type.GetValue(); !ok {
			logus.Log.Warn("ship problem with type", typelog.Any("nickname", ship.Nickname))
		}
		ship.Type, _ = ship_info.Type.GetValue()
		ship.Type = strings.ToLower(ship.Type)

		if ship_name_id, ship_has_name := ship_info.IdsName.GetValue(); ship_has_name {
			ship.NameID = ship_name_id
		} else {
			logus.Log.Warn("WARNING, ship has no ItdsName", typelog.String("ship.Nickname", ship.Nickname))
		}

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

		ship.Name = e.GetInfocardName(ship.NameID, ship.Nickname)

		if ship_hull_good, ok := e.configs.Goods.ShipHullsMapByShip[ship.Nickname]; ok {
			ship.Price = ship_hull_good.Price.Get()

			ship_hull_nickname := ship_hull_good.Nickname.Get()
			if ship_package_goods, ok := e.configs.Goods.ShipsMapByHull[ship_hull_nickname]; ok {

				for _, ship_package_good := range ship_package_goods {
					for _, addon := range ship_package_good.Addons {

						// can be Power or Engine or Smth else
						// addon = dsy_hessian_engine, HpEngine01, 1
						// addon = dsy_loki_core, internal, 1
						// addon = ge_s_scanner_01, internal, 1
						addon_nickname := addon.ItemNickname.Get()

						if good_info, ok := e.configs.Goods.GoodsMap[addon_nickname]; ok {
							if addon_price, ok := good_info.Price.GetValue(); ok {
								ship.Price += addon_price
							}
						}

						if power, ok := e.configs.Equip.PowersMap[addon_nickname]; ok {
							ship.PowerCapacity = power.Capacity.Get()
							ship.PowerRechargeRate = power.ChargeRate.Get()

							ship.ThrustCapacity = power.ThrustCapacity.Get()
							ship.ThrustRecharge = power.ThrustRecharge.Get()
						}
						if engine, ok := e.configs.Equip.EnginesMap[addon_nickname]; ok {
							ship.CruiseSpeed = e.GetEngineSpeed(engine)
							engine_linear_drag, _ := engine.LinearDrag.GetValue()
							ship_linear_drag, _ := ship_info.LinearDrag.GetValue()
							engine_max_force, _ := engine.MaxForce.GetValue()
							ship.ImpulseSpeed = float64(engine_max_force) / (float64(engine_linear_drag) + float64(ship_linear_drag))

							ship.ReverseFraction = engine.ReverseFraction.Get()

							ship.MaxAngularSpeedDegS = ship_info.SteeringTorque.X.Get() / ship_info.AngularDrag.X.Get()
							ship.TimeTo90MaxAngularSpeed = ship_info.RotationIntertia.X.Get() / (ship_info.AngularDrag.X.Get() * LogOgE)

							ship.MaxAngularSpeedDegS *= Pi180

							if ship.TimeTo90MaxAngularSpeed > 0.5 {
								ship.AngularDistanceFrom0ToHalfSec = ship.MaxAngularSpeedDegS * (0.5 / ship.TimeTo90MaxAngularSpeed) / 2
							} else {
								ship.AngularDistanceFrom0ToHalfSec = ship.MaxAngularSpeedDegS*(0.5-ship.TimeTo90MaxAngularSpeed) + ship.MaxAngularSpeedDegS*ship.TimeTo90MaxAngularSpeed/2
							}
						}
					}

					ships_at_bases := e.GetAtBasesSold(GetCommodityAtBasesInput{
						Nickname: ship_package_good.Nickname.Get(),
						Price:    ship.Price,
					})
					for key, value := range ships_at_bases {
						ship.Bases[key] = value
					}
				}

			}

		}

		ship.HoldSize = ship_info.HoldSize.Get()
		ship.Armor = ship_info.HitPts.Get()

		var hardpoints map[string][]string = make(map[string][]string)
		for _, hp_type := range ship_info.HpTypes {
			for _, equipment := range hp_type.AllowedEquipments {
				equipment_slot := equipment.Get()
				hardpoints[equipment_slot] = append(hardpoints[equipment_slot], hp_type.Nickname.Get())
			}
		}

		for slot_name, allowed_equip := range hardpoints {
			ship.Slots = append(ship.Slots, EquipmentSlot{
				SlotName:     slot_name,
				AllowedEquip: allowed_equip,
			})
		}

		sort.Slice(ship.Slots, func(i, j int) bool {
			return ship.Slots[i].SlotName < ship.Slots[j].SlotName
		})
		for _, slot := range ship.Slots {
			sort.Slice(slot.AllowedEquip, func(i, j int) bool {
				return slot.AllowedEquip[i] < slot.AllowedEquip[j]
			})
		}

		for _, slot := range ship.Slots {
			if len(slot.AllowedEquip) > len(ship.BiggestHardpoint) {
				ship.BiggestHardpoint = slot.AllowedEquip
			}
		}

		var infocards []int
		if id, ok := ship_info.IdsInfo1.GetValue(); ok {
			infocards = append(infocards, id)
		}
		// if id, ok := ship_info.IdsInfo2.GetValue(); ok {
		// 	infocards = append(infocards, id)
		// }
		// Nobody uses it?
		// if id, ok := ship_info.IdsInfo3.GetValue(); ok {
		// 	infocards = append(infocards, id)
		// }
		if id, ok := ship_info.IdsInfo.GetValue(); ok {
			infocards = append(infocards, id)
		}
		e.exportInfocards(InfocardKey(ship.Nickname), infocards...)
		ship.DiscoveryTechCompat = CalculateTechCompat(e.configs.Discovery, ids, ship.Nickname)

		if e.configs.Discovery != nil {
			armor_mult, _ := ship_info.ArmorMult.GetValue()
			ship.DiscoShip = &DiscoShip{ArmorMult: armor_mult}
		}

		ships = append(ships, ship)
	}

	return ships
}

type EquipmentSlot struct {
	SlotName     string
	AllowedEquip []string
}

var Pi180 = 180 / math.Pi // number turning radians to degrees
var LogOgE = math.Log10(math.E)

func (e *Exporter) FilterToUsefulShips(ships []Ship) []Ship {
	var items []Ship = make([]Ship, 0, len(ships))
	for _, item := range ships {
		if !e.Buyable(item.Bases) {
			continue
		}
		items = append(items, item)
	}
	return items
}

type CompatibleIDsForTractor struct {
	TechCompat float64
	Tractor    Tractor
}

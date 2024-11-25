package configs_export

import (
	"sync"

	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_export/trades"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld/flhash"
	"github.com/darklab8/fl-configs/configs/configs_settings"
)

type InfocardKey string

type Infocard []string

func (e *Exporter) exportInfocards(nickname InfocardKey, infocard_ids ...int) {
	if _, ok := e.Infocards[InfocardKey(nickname)]; ok {
		return
	}

	for _, info_id := range infocard_ids {
		if value, ok := e.configs.Infocards.Infocards[info_id]; ok {
			e.Infocards[InfocardKey(nickname)] = append(e.Infocards[InfocardKey(nickname)], value.Lines...)
		}
	}

	if len(e.Infocards[InfocardKey(nickname)]) == 0 {
		e.Infocards[InfocardKey(nickname)] = []string{"no infocard"}
	}
}

type Infocards map[InfocardKey]Infocard

type Exporter struct {
	configs *configs_mapped.MappedConfigs
	Hashes  map[string]flhash.HashCode

	Bases                []*Base
	MiningOperations     []*Base
	useful_bases_by_nick map[cfgtype.BaseUniNick]bool

	ship_speeds trades.ShipSpeeds
	transport   *GraphResults
	freighter   *GraphResults
	frigate     *GraphResults

	Factions     []Faction
	Infocards    Infocards
	Commodities  []*Commodity
	Guns         []Gun
	Missiles     []Gun
	Mines        []Mine
	Shields      []Shield
	Thrusters    []Thruster
	Ships        []Ship
	Tractors     []Tractor
	TractorsByID map[cfgtype.TractorID]Tractor
	Engines      []Engine
	CMs          []CounterMeasure
	Scanners     []Scanner
	Ammos        []Ammo
}

type OptExport func(e *Exporter)

func NewExporter(configs *configs_mapped.MappedConfigs, opts ...OptExport) *Exporter {
	e := &Exporter{
		configs:     configs,
		Infocards:   map[InfocardKey]Infocard{},
		ship_speeds: trades.VanillaSpeeds,
		Hashes:      make(map[string]flhash.HashCode),
	}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

type GraphResults struct {
	e       *Exporter
	graph   *trades.GameGraph
	dists   [][]int
	parents [][]trades.Parent
}

func NewGraphResults(
	e *Exporter,
	avgCruiserSpeed int,
	can_visit_freighter_only_jhs trades.WithFreighterPaths,
	mining_bases_by_system map[string][]trades.ExtraBase,
) *GraphResults {
	graph := trades.MapConfigsToFGraph(e.configs, avgCruiserSpeed, can_visit_freighter_only_jhs, mining_bases_by_system)
	dijkstra_apsp := trades.NewDijkstraApspFromGraph(graph)
	dists, parents := dijkstra_apsp.DijkstraApsp()

	return &GraphResults{
		e:       e,
		graph:   graph,
		dists:   dists,
		parents: parents,
	}
}

func (e *Exporter) Export() *Exporter {
	var wg sync.WaitGroup

	e.Bases = e.GetBases()
	useful_bases := FilterToUserfulBases(e.Bases)
	e.useful_bases_by_nick = make(map[cfgtype.BaseUniNick]bool)
	for _, base := range useful_bases {
		e.useful_bases_by_nick[base.Nickname] = true
	}
	e.useful_bases_by_nick[pob_crafts_nickname] = true

	e.Commodities = e.GetCommodities()
	EnhanceBasesWithServerOverrides(e.Bases, e.Commodities)

	e.MiningOperations = e.GetOres(e.Commodities)
	mining_bases_by_system := make(map[string][]trades.ExtraBase)
	for _, base := range e.MiningOperations {
		mining_bases_by_system[base.SystemNickname] = append(mining_bases_by_system[base.SystemNickname], trades.ExtraBase{
			Pos:      base.Pos,
			Nickname: base.Nickname,
		})
	}
	if e.configs.Discovery != nil {
		e.ship_speeds = trades.DiscoverySpeeds
	}

	if !configs_settings.Env.IsDisabledTradeRouting {
		wg.Add(1)
		go func() {
			e.transport = NewGraphResults(e, e.ship_speeds.AvgTransportCruiseSpeed, trades.WithFreighterPaths(false), mining_bases_by_system)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			e.freighter = NewGraphResults(e, e.ship_speeds.AvgFreighterCruiseSpeed, trades.WithFreighterPaths(true), mining_bases_by_system)
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			e.frigate = NewGraphResults(e, e.ship_speeds.AvgFrigateCruiseSpeed, trades.WithFreighterPaths(false), mining_bases_by_system)
			wg.Done()
		}()
	}

	e.Tractors = e.GetTractors()
	e.TractorsByID = make(map[cfgtype.TractorID]Tractor)
	for _, tractor := range e.Tractors {
		e.TractorsByID[tractor.Nickname] = tractor
	}
	e.Factions = e.GetFactions(e.Bases)
	e.Bases = e.GetMissions(e.Bases, e.Factions)

	e.Shields = e.GetShields(e.Tractors)
	buyable_shield_tech := e.GetBuyableShields(e.Shields)
	e.Guns = e.GetGuns(e.Tractors, buyable_shield_tech)
	e.Missiles = e.GetMissiles(e.Tractors, buyable_shield_tech)
	e.Mines = e.GetMines(e.Tractors)

	e.Thrusters = e.GetThrusters(e.Tractors)
	e.Ships = e.GetShips(e.Tractors, e.TractorsByID)
	e.Engines = e.GetEngines(e.Tractors)
	e.CMs = e.GetCounterMeasures(e.Tractors)
	e.Scanners = e.GetScanners(e.Tractors)
	e.Ammos = e.GetAmmo(e.Tractors)
	wg.Wait()

	e.Bases, e.Commodities = e.TradePaths(e.Bases, e.Commodities)
	e.MiningOperations, e.Commodities = e.TradePaths(e.MiningOperations, e.Commodities)
	e.Bases = e.AllRoutes(e.Bases)

	for _, system := range e.configs.Systems.Systems {
		for zone_nick := range system.ZonesByNick {
			e.Hashes[zone_nick] = flhash.HashNickname(zone_nick)
		}
		for _, object := range system.Objects {
			nickname, _ := object.Nickname.GetValue()
			e.Hashes[nickname] = flhash.HashNickname(nickname)
		}
	}
	for _, good := range e.configs.Goods.Goods {
		nickname, _ := good.Nickname.GetValue()
		e.Hashes[nickname] = flhash.HashNickname(nickname)
	}

	e.EnhanceBasesWithIsTransportReachable(e.Bases, e.transport)
	e.Bases = e.EnhanceBasesWithPobCrafts(e.Bases)

	return e
}

func (e *Exporter) EnhanceBasesWithIsTransportReachable(
	bases []*Base,
	transports_graph *GraphResults,
) {
	reachable_base_example := "li01_01_base"
	g := transports_graph

	for _, base := range bases {
		base_nickname := base.Nickname.ToStr()
		if trades.GetDist(g.graph, g.dists, reachable_base_example, base_nickname) >= trades.INF/2 {
			base.IsTransportUnreachable = true
		}
	}

	enhance_with_transport_unrechability := func(Bases map[cfgtype.BaseUniNick]*GoodAtBase) {
		for _, base := range Bases {
			if trades.GetDist(g.graph, g.dists, reachable_base_example, string(base.BaseNickname)) >= trades.INF/2 {
				base.IsTransportUnreachable = true
			}
		}
	}

	for _, item := range e.Commodities {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Guns {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Missiles {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Mines {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Shields {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Thrusters {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Ships {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Tractors {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Engines {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.CMs {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Scanners {
		enhance_with_transport_unrechability(item.Bases)
	}
	for _, item := range e.Ammos {
		enhance_with_transport_unrechability(item.Bases)
	}
}

func Export(configs *configs_mapped.MappedConfigs) *Exporter {
	return NewExporter(configs).Export()
}

func Empty(phrase string) bool {
	for _, letter := range phrase {
		if letter != ' ' {
			return false
		}
	}
	return true
}

func (e *Exporter) Buyable(Bases map[cfgtype.BaseUniNick]*GoodAtBase) bool {
	for _, base := range Bases {

		if e.useful_bases_by_nick != nil {
			if _, ok := e.useful_bases_by_nick[base.BaseNickname]; ok {
				return true
			}
		}
	}

	return false
}

func Buyable(Bases map[cfgtype.BaseUniNick]*GoodAtBase) bool {
	return len(Bases) > 0
}

type DiscoveryTechCompat struct {
	TechcompatByID map[cfgtype.TractorID]float64
	TechCell       string
}

func CalculateTechCompat(Discovery *configs_mapped.DiscoveryConfig, ids []Tractor, nickname string) *DiscoveryTechCompat {
	if Discovery == nil {
		return nil
	}

	techcompat := &DiscoveryTechCompat{
		TechcompatByID: make(map[cfgtype.TractorID]float64),
	}
	techcompat.TechcompatByID[""] = Discovery.Techcompat.GetCompatibilty(nickname, "")

	for _, id := range ids {
		techcompat.TechcompatByID[id.Nickname] = Discovery.Techcompat.GetCompatibilty(nickname, id.Nickname)
	}

	if compat, ok := Discovery.Techcompat.CompatByItem[nickname]; ok {
		techcompat.TechCell = compat.TechCell
	}

	return techcompat
}

func (e *Exporter) GetInfocardName(ids_name int, nickname string) string {
	return e.configs.GetInfocardName(ids_name, nickname)
}

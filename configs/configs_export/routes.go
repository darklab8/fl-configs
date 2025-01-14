package configs_export

import (
	"github.com/darklab8/fl-configs/configs/cfgtype"
	"github.com/darklab8/fl-configs/configs/configs_export/trades"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
)

type Route struct {
	g                  *GraphResults
	is_disabled        bool
	from_base_nickname string
	to_base_nickname   string
}

type BaseRoute struct {
	*Route
	FromBase *Base
	ToBase   *Base
}

func NewBaseRoute(g *GraphResults, FromBase *Base, ToBase *Base) *BaseRoute {
	return &BaseRoute{
		Route:    NewRoute(g, FromBase.Nickname.ToStr(), ToBase.Nickname.ToStr()),
		FromBase: FromBase,
		ToBase:   ToBase,
	}

}

type BaseAllRoutes struct {
	AllRoutes []*ComboRoute
}

type ComboRoute struct {
	Transport *BaseRoute
	Frigate   *BaseRoute
	Freighter *BaseRoute
}

func NewRoute(g *GraphResults, from_base_nickname string, to_base_nickname string) *Route {
	return &Route{
		g:                  g,
		from_base_nickname: from_base_nickname,
		to_base_nickname:   to_base_nickname,
	}
}

func (c *Route) GetID() string {
	if c.is_disabled {
		return ""
	}
	return c.from_base_nickname + c.to_base_nickname
}

func (t *Route) GetCruiseSpeed() int {
	if t.is_disabled {
		return 0
	}
	return t.g.graph.AvgCruiseSpeed
}

func (t *Route) GetCanVisitFreighterOnlyJH() bool {
	if t.is_disabled {
		return false
	}
	return bool(t.g.graph.CanVisitFreightersOnlyJHs)
}

type PathWithNavmap struct {
	trades.DetailedPath
	SectorCoord string
	Pos         cfgtype.Vector
}

func (t *Route) GetPaths() []PathWithNavmap {
	var results []PathWithNavmap
	paths := t.g.graph.GetPaths(t.g.parents, t.g.dists, t.from_base_nickname, t.to_base_nickname)

	for _, path := range paths {
		// path.NextName // nickname of object

		augmented_path := PathWithNavmap{
			DetailedPath: path,
		}

		if jh, ok := t.g.e.Configs.Systems.JumpholesByNick[path.NextName]; ok {
			pos := jh.Pos.Get()

			system_uni := t.g.e.Configs.Universe.SystemMap[universe_mapped.SystemNickname(jh.System.Nickname)]
			augmented_path.SectorCoord = VectorToSectorCoord(system_uni, pos)
			augmented_path.Pos = pos
		}
		if base, ok := t.g.e.Configs.Systems.BasesByDockWith[path.NextName]; ok {
			pos := base.Pos.Get()

			system_uni := t.g.e.Configs.Universe.SystemMap[universe_mapped.SystemNickname(base.System.Nickname)]
			augmented_path.SectorCoord = VectorToSectorCoord(system_uni, pos)
			augmented_path.Pos = pos
		}

		results = append(results, augmented_path)
	}
	return results
}

func (t *Route) GetNameByIdsName(ids_name int) string {
	return string(t.g.e.Configs.Infocards.Infonames[ids_name])
}

func (t *Route) GetDist() int {
	return trades.GetDist(t.g.graph, t.g.dists, t.from_base_nickname, t.to_base_nickname)
}

func (t *Route) GetTime() float64 {
	return float64(t.GetDist())/trades.PrecisionMultipiler + float64(trades.BaseDockingDelay)
}

func (e *Exporter) AllRoutes(
	bases []*Base,
) []*Base {
	for _, from_base := range bases {
		for _, to_base := range bases {
			// it can fly everywhere so we use it for checking
			freighter_route := NewBaseRoute(e.freighter, from_base, to_base)

			if freighter_route.GetDist() > trades.INF/2 {
				continue
			}

			from_base.AllRoutes = append(from_base.AllRoutes, &ComboRoute{
				Transport: NewBaseRoute(e.transport, from_base, to_base),
				Frigate:   NewBaseRoute(e.frigate, from_base, to_base),
				Freighter: freighter_route,
			})
		}
	}
	return bases
}

package configs_export

import "github.com/darklab8/fl-configs/configs/configs_mapped"

type Exporter struct {
	configs                    *configs_mapped.MappedConfigs
	are_no_name_bases_included NoNameIncluded

	Bases         []Base
	GoodsSelEquip []GoodSelEquip
	Factions      []Faction
}

type OptExport func(e *Exporter)

func WithNoNameBases() OptExport {
	return func(e *Exporter) { e.are_no_name_bases_included = true }
}

func NewExporter(configs *configs_mapped.MappedConfigs, opts ...OptExport) *Exporter {
	e := &Exporter{
		configs:                    configs,
		are_no_name_bases_included: false,
	}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Exporter) Export() *Exporter {
	e.Bases = e.getBases(e.are_no_name_bases_included)
	e.GoodsSelEquip = e.getGoodSelEquip()
	e.Factions = e.GetFactions()
	return e
}

func Export(configs *configs_mapped.MappedConfigs) *Exporter {
	return NewExporter(configs).Export()
}

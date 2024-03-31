package configs_export

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped"
)

type Exporter struct {
	configs                    *configs_mapped.MappedConfigs
	are_no_name_bases_included NoNameIncluded

	Bases       []Base
	Factions    []Faction
	Infocards   map[InfocardKey]*Infocard
	Commodities []Commodity
	Guns        []Gun

	infocards_parser *InfocardsParser
}

type OptExport func(e *Exporter)

func WithNoNameBases() OptExport {
	return func(e *Exporter) { e.are_no_name_bases_included = true }
}

func NewExporter(configs *configs_mapped.MappedConfigs, opts ...OptExport) *Exporter {
	e := &Exporter{
		configs:                    configs,
		are_no_name_bases_included: false,
		infocards_parser:           NewInfocardsParser(configs.Infocards),
	}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Exporter) Export() *Exporter {
	e.Bases = e.GetBases(e.are_no_name_bases_included)
	e.Factions = e.GetFactions(e.Bases)
	e.Commodities = e.GetCommodities()
	e.Guns = e.GetGuns()
	e.Infocards = e.infocards_parser.Get()
	return e
}

func Export(configs *configs_mapped.MappedConfigs) *Exporter {
	return NewExporter(configs).Export()
}

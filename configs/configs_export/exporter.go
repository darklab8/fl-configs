package configs_export

import "github.com/darklab8/fl-configs/configs/configs_mapped"

type Exporter struct {
	configs             *configs_mapped.MappedConfigs
	is_no_name_included NoNameIncluded

	Bases         []Base
	GoodsSelEquip []GoodSelEquip
}

type OptExport func(e *Exporter)

func WithNonameIncluded() OptExport {
	return func(e *Exporter) { e.is_no_name_included = false }
}

func NewExporter(configs *configs_mapped.MappedConfigs, opts ...OptExport) *Exporter {
	e := &Exporter{
		configs:             configs,
		is_no_name_included: true,
	}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *Exporter) Export() *Exporter {
	e.Bases = e.getBases(e.is_no_name_included)
	e.GoodsSelEquip = e.getGoodSelEquip()
	return e
}

func Export(configs *configs_mapped.MappedConfigs) *Exporter {
	return NewExporter(configs).Export()
}

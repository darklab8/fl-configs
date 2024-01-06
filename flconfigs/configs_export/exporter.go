package configs_export

import "github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped"

type Exporter struct {
	configs *configs_mapped.MappedConfigs
}

func NewExporter(configs *configs_mapped.MappedConfigs) *Exporter {
	e := &Exporter{
		configs: configs,
	}
	return e
}

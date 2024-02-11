package configs_export

import (
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func (e *Exporter) Bases() []Base {
	var results []Base = make([]Base, len(e.configs.Universe_config.Bases))
	for number, base := range e.configs.Universe_config.Bases {
		results[number] = Base{
			Nickname:         base.Nickname.Get(),
			System:           base.System.Get(),
			StridName:        base.StridName.Get(),
			File:             utils_types.FilePath(base.File.Get()),
			BGCS_base_run_by: base.BGCS_base_run_by.Get(),
		}
	}
	return results
}

type Base struct {
	Nickname         string
	System           string
	StridName        int
	File             utils_types.FilePath
	BGCS_base_run_by string
}

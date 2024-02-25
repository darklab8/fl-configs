package configs_export

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type NoNameIncluded bool

func (e *Exporter) Bases(is_no_name_included NoNameIncluded) []Base {
	var results []Base = make([]Base, len(e.configs.Universe_config.Bases))

	iterator := 0
	for _, base := range e.configs.Universe_config.Bases {

		var name string
		if base_infocard, ok := e.configs.Infocards.RecordsMap[base.StridName.Get()]; ok {
			name = base_infocard.Content
		}

		var system_name string
		if system, ok := e.configs.Universe_config.SystemMap.MapGetValue(universe_mapped.SystemNickname(base.System.Get())); ok {

			if infocard, ok := e.configs.Infocards.RecordsMap[system.Strid_name.Get()]; ok {
				system_name = infocard.Content
			}

		}

		if !is_no_name_included && name == "" {
			continue
		}

		results[iterator] = Base{
			Name:             name,
			Nickname:         base.Nickname.Get(),
			System:           system_name,
			SystemNickname:   base.System.Get(),
			StridName:        base.StridName.Get(),
			File:             utils_types.FilePath(base.File.Get()),
			BGCS_base_run_by: base.BGCS_base_run_by.Get(),
		}
		iterator += 1
	}

	results = results[:iterator]
	return results
}

type Base struct {
	Name             string
	Nickname         string
	System           string
	SystemNickname   string
	StridName        int
	File             utils_types.FilePath
	BGCS_base_run_by string
}

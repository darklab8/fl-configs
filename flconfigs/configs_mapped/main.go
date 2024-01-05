/*
Tool to parse freelancer configs
*/
package configs_mapped

import (
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/data/universe_mapped"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/data/universe_mapped/systems_mapped"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_flconfigs/flconfigs/settings/logus"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/freelancer_mapped/data/equipment_mapped/market_mapped"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type MappedConfigs struct {
	Universe_config *universe_mapped.Config
	// Info_config         *infocard.Config
	Systems             *systems_mapped.Config
	Market_ships_config *market_mapped.Config
	Market_commodities  *market_mapped.Config
	Market_misc         *market_mapped.Config
}

func NewMappedConfigs() *MappedConfigs {
	return &MappedConfigs{}
}

func (p *MappedConfigs) Read(file1path utils_types.FilePath) *MappedConfigs {
	logus.Log.Info("Parse START for FreelancerFolderLocation=", logus_core.FilePath(file1path))
	filesystem := filefind.FindConfigs(file1path)

	p.Universe_config = (&universe_mapped.Config{}).Read(filesystem.GetFile(universe_mapped.FILENAME))
	// p.Info_config = (&infocard.Config{}).Read(filesystem.GetFile(infocard.FILENAME, infocard.FILENAME_FALLBACK))
	p.Systems = (&systems_mapped.Config{}).Read(p.Universe_config, filesystem)
	p.Market_ships_config = (&market_mapped.Config{}).Read(filesystem.GetFile(market_mapped.FILENAME_SHIPS))
	p.Market_commodities = (&market_mapped.Config{}).Read(filesystem.GetFile(market_mapped.FILENAME_COMMODITIES))
	p.Market_misc = (&market_mapped.Config{}).Read(filesystem.GetFile(market_mapped.FILENAME_MISC))

	logus.Log.Info("Parse OK for FreelancerFolderLocation=", logus_core.FilePath(file1path))

	return p
}

type IsDruRun bool

func (p *MappedConfigs) Write(is_dry_run IsDruRun) {
	// write
	files := []*file.File{}

	files = append(files,
		p.Universe_config.Write(),
		p.Market_ships_config.Write(),
		p.Market_commodities.Write(),
		p.Market_misc.Write(),
	)
	files = append(files, p.Systems.Write()...)

	if is_dry_run {
		return
	}

	for _, file := range files {
		file.WriteLines()
	}
}

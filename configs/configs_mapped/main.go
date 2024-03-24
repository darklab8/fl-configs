/*
Tool to parse freelancer configs
*/
package configs_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped/systems_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/settings/logus"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/market_mapped"

	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type MappedConfigs struct {
	Universe_config   *universe_mapped.Config
	Systems           *systems_mapped.Config
	MarketCapital     *market_mapped.Config
	MarketShips       *market_mapped.Config
	MarketCommidities *market_mapped.Config
	MarketMisc        *market_mapped.Config
	FreelancerINI     *exe_mapped.Config
	Infocards         *infocard.Config
}

func NewMappedConfigs() *MappedConfigs {
	return &MappedConfigs{}
}

func (p *MappedConfigs) Read(file1path utils_types.FilePath) *MappedConfigs {
	logus.Log.Info("Parse START for FreelancerFolderLocation=", utils_logus.FilePath(file1path))
	filesystem := filefind.FindConfigs(file1path)

	p.Universe_config = universe_mapped.Read(filesystem.GetFile(universe_mapped.FILENAME))
	p.Systems = systems_mapped.Read(p.Universe_config, filesystem)

	p.MarketCapital = market_mapped.Read(filesystem.GetFile("market_capital.ini"))
	p.MarketCommidities = market_mapped.Read(filesystem.GetFile(market_mapped.FILENAME_COMMODITIES))
	p.MarketMisc = market_mapped.Read(filesystem.GetFile(market_mapped.FILENAME_MISC))
	p.MarketShips = market_mapped.Read(filesystem.GetFile(market_mapped.FILENAME_SHIPS))
	p.FreelancerINI = exe_mapped.Read(filesystem.GetFile(exe_mapped.FILENAME_FL_INI))

	p.Infocards = infocard_mapped.Read(filesystem, p.FreelancerINI, filesystem.GetFile(infocard_mapped.FILENAME, infocard_mapped.FILENAME_FALLBACK))

	logus.Log.Info("Parse OK for FreelancerFolderLocation=", utils_logus.FilePath(file1path))

	return p
}

type IsDruRun bool

func (p *MappedConfigs) Write(is_dry_run IsDruRun) {
	// write
	files := []*file.File{}

	files = append(files,
		p.Universe_config.Write(),
		p.MarketShips.Write(),
		p.MarketCommidities.Write(),
		p.MarketMisc.Write(),
	)
	files = append(files, p.Systems.Write()...)

	if is_dry_run {
		return
	}

	for _, file := range files {
		file.WriteLines()
	}
}

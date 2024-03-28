/*
Tool to parse freelancer configs
*/
package configs_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/initialworld"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/interface_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/missions_mapped/empathy_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/ship_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped/systems_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/settings/logus"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/equip_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/equipment_mapped/market_mapped"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type MappedConfigs struct {
	FreelancerINI *exe_mapped.Config

	Universe_config *universe_mapped.Config
	Systems         *systems_mapped.Config

	Market   *market_mapped.Config
	Equip    *equip_mapped.Config
	Goods    *equipment_mapped.Config
	Shiparch *ship_mapped.Config

	InfocardmapINI *interface_mapped.Config
	Infocards      *infocard.Config
	InitialWorld   *initialworld.Config
	Empathy        *empathy_mapped.Config
}

func NewMappedConfigs() *MappedConfigs {
	return &MappedConfigs{}
}

func (p *MappedConfigs) Read(file1path utils_types.FilePath) *MappedConfigs {
	logus.Log.Info("Parse START for FreelancerFolderLocation=", utils_logus.FilePath(file1path))
	filesystem := filefind.FindConfigs(file1path)

	p.FreelancerINI = exe_mapped.Read(filesystem.GetFile(exe_mapped.FILENAME_FL_INI))
	p.Universe_config = universe_mapped.Read(filesystem.GetFile(universe_mapped.FILENAME))
	p.Systems = systems_mapped.Read(p.Universe_config, filesystem)

	get_files := func(paths []*semantic.Path) []*file.File {
		return utils.CompL(paths, func(x *semantic.Path) *file.File { return filesystem.GetFile(utils_types.FilePath(x.FileName())) })
	}

	p.Market = market_mapped.Read(get_files(p.FreelancerINI.Markets))
	p.Equip = equip_mapped.Read(get_files(p.FreelancerINI.Equips))
	p.Goods = equipment_mapped.Read(get_files(p.FreelancerINI.Goods))
	p.Shiparch = ship_mapped.Read(get_files(p.FreelancerINI.Ships))

	p.InfocardmapINI = interface_mapped.Read(filesystem.GetFile(interface_mapped.FILENAME_FL_INI))
	p.Infocards = infocard_mapped.Read(filesystem, p.FreelancerINI, filesystem.GetFile(infocard_mapped.FILENAME, infocard_mapped.FILENAME_FALLBACK))

	p.InitialWorld = initialworld.Read(filesystem.GetFile(initialworld.FILENAME))
	p.Empathy = empathy_mapped.Read(filesystem.GetFile(empathy_mapped.FILENAME))

	logus.Log.Info("Parse OK for FreelancerFolderLocation=", utils_logus.FilePath(file1path))

	return p
}

type IsDruRun bool

func (p *MappedConfigs) Write(is_dry_run IsDruRun) {
	// write
	files := []*file.File{}

	files = append(files,
		p.Universe_config.Write(),
	)
	files = append(files, p.Systems.Write()...)
	files = append(files, p.Market.Write()...)

	if is_dry_run {
		return
	}

	for _, file := range files {
		file.WriteLines()
	}
}

/*
Tool to parse freelancer configs
*/
package parser

import (
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/freelancer/data/universe"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/freelancer/data/universe/systems"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/parserutils/filefind"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/parserutils/filefind/file"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/parser/freelancer/data/equipment/market"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type Parsed struct {
	Universe_config *universe.Config
	// Info_config         *infocard.Config
	Systems             *systems.Config
	Market_ships_config *market.Config
	Market_commodities  *market.Config
	Market_misc         *market.Config
}

func NewParsed() *Parsed {
	return &Parsed{}
}

func (p *Parsed) Read(file1path utils_types.FilePath) *Parsed {
	logus.Log.Info("Parse START for FreelancerFolderLocation=", logus_core.FilePath(file1path))
	filesystem := filefind.FindConfigs(file1path)

	p.Universe_config = (&universe.Config{}).Read(filesystem.GetFile(universe.FILENAME))
	// p.Info_config = (&infocard.Config{}).Read(filesystem.GetFile(infocard.FILENAME, infocard.FILENAME_FALLBACK))
	p.Systems = (&systems.Config{}).Read(p.Universe_config, filesystem)
	p.Market_ships_config = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_SHIPS))
	p.Market_commodities = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_COMMODITIES))
	p.Market_misc = (&market.Config{}).Read(filesystem.GetFile(market.FILENAME_MISC))

	logus.Log.Info("Parse OK for FreelancerFolderLocation=", logus_core.FilePath(file1path))

	return p
}

type IsDruRun bool

func (p *Parsed) Write(is_dry_run IsDruRun) {
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

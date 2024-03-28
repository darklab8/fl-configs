/*
parse universe.ini
*/
package universe_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
)

// Feel free to map it xD
// terrain_tiny = nonmineable_asteroid90
// terrain_sml = nonmineable_asteroid60
// terrain_mdm = nonmineable_asteroid90
// terrain_lrg = nonmineable_asteroid60
// terrain_dyna_01 = mineable1_asteroid10
// terrain_dyna_02 = mineable1_asteroid10

var KEY_BASE_TERRAINS = [...]string{"terrain_tiny", "terrain_sml", "terrain_mdm", "terrain_lrg", "terrain_dyna_01", "terrain_dyna_02"}

const (
	FILENAME      = "universe.ini"
	KEY_BASE_TAG  = "[Base]"
	KEY_NICKNAME  = "nickname"
	KEY_STRIDNAME = "strid_name"
	KEY_SYSTEM    = "system"
	KEY_FILE      = "file"

	KEY_BASE_BGCS = "BGCS_base_run_by"

	KEY_SYSTEM_TAG           = "[system]"
	KEY_SYSTEM_MSG_ID_PREFIX = "msg_id_prefix"
	KEY_SYSTEM_VISIT         = "visit"
	KEY_SYSTEM_IDS_INFO      = "ids_info"
	KEY_SYSTEM_NAVMAPSCALE   = "NavMapScale"
	KEY_SYSTEM_POS           = "pos"

	KEY_TIME_TAG     = "[Time]"
	KEY_TIME_SECONDS = "seconds_per_day"
)

type Base struct {
	semantic.Model

	Nickname         *semantic.String
	System           *semantic.String
	StridName        *semantic.Int
	File             *semantic.Path
	BGCS_base_run_by *semantic.String
	// Terrains *semantic.StringStringMap
}

type BaseNickname string

type SystemNickname string

type System struct {
	semantic.Model
	Nickname *semantic.String
	// Pos        *semantic.Pos
	Msg_id_prefix *semantic.String
	Visit         *semantic.Int
	Strid_name    *semantic.Int
	Ids_info      *semantic.Int
	File          *semantic.Path
	// NavMapScale   *semantic.Float
}

type Config struct {
	semantic.ConfigModel
	Bases    []*Base
	BasesMap *lower_map.KeyLoweredMap[BaseNickname, *Base]

	Systems   []*System
	SystemMap *lower_map.KeyLoweredMap[SystemNickname, *System]

	TimeSeconds *semantic.Int
}

func Read(input_file *file.File) *Config {
	frelconfig := &Config{}

	iniconfig := inireader.Read(input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	frelconfig.TimeSeconds = semantic.NewInt(iniconfig.SectionMap[KEY_TIME_TAG][0], KEY_TIME_TAG)
	frelconfig.BasesMap = lower_map.NewKeyLoweredMap[BaseNickname, *Base]()
	frelconfig.Bases = make([]*Base, 0)
	frelconfig.SystemMap = lower_map.NewKeyLoweredMap[SystemNickname, *System]()
	frelconfig.Systems = make([]*System, 0)

	if bases, ok := iniconfig.SectionMap[KEY_BASE_TAG]; ok {
		for _, base := range bases {
			base_to_add := &Base{}
			base_to_add.Map(base)
			base_to_add.Nickname = semantic.NewString(base, KEY_NICKNAME)
			base_to_add.StridName = semantic.NewInt(base, KEY_STRIDNAME)
			base_to_add.BGCS_base_run_by = semantic.NewString(base, KEY_BASE_BGCS, semantic.OptsS(semantic.Optional()))
			base_to_add.System = semantic.NewString(base, KEY_SYSTEM)
			base_to_add.File = semantic.NewPath(base, KEY_FILE)

			frelconfig.Bases = append(frelconfig.Bases, base_to_add)
			frelconfig.BasesMap.MapSet(BaseNickname(base_to_add.Nickname.Get()), base_to_add)
		}
	}

	if systems, ok := iniconfig.SectionMap[KEY_SYSTEM_TAG]; ok {
		for _, system := range systems {
			system_to_add := System{}
			system_to_add.Map(system)

			system_to_add.Visit = semantic.NewInt(system, KEY_SYSTEM_VISIT, semantic.Optional())
			system_to_add.Strid_name = semantic.NewInt(system, KEY_STRIDNAME, semantic.Optional())
			system_to_add.Ids_info = semantic.NewInt(system, KEY_SYSTEM_IDS_INFO, semantic.Optional())
			// system_to_add.NavMapScale = system.GetParamNumber(KEY_SYSTEM_NAVMAPSCALE, inireader.OPTIONAL_p)
			system_to_add.Nickname = semantic.NewString(system, KEY_NICKNAME)
			system_to_add.File = semantic.NewPath(system, KEY_FILE)
			system_to_add.Msg_id_prefix = semantic.NewString(system, KEY_SYSTEM_MSG_ID_PREFIX, semantic.OptsS(semantic.Optional()))

			frelconfig.Systems = append(frelconfig.Systems, &system_to_add)
			frelconfig.SystemMap.MapSet(SystemNickname(system_to_add.Nickname.Get()), &system_to_add)
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

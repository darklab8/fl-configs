/*
parse universe.ini
*/
package universe_mapped

import (
	"strings"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/semantic"
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

	KEY_NAME    = "name"
	KEY_RECYCLE = "is_recycle_candidate"
)

type Base struct {
	semantic.Model

	Nickname         *semantic.String
	System           *semantic.String
	StridName        *semantic.Int
	File             *semantic.Path
	BGCS_base_run_by *semantic.String
	// Terrains *semantic.StringStringMap

	Name             *semantic.String // denormalized always disabled param
	RecycleCandidate *semantic.String // denormalized always disabled param
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
	BasesMap map[BaseNickname]*Base

	Systems   []*System
	SystemMap map[SystemNickname]*System

	TimeSeconds *semantic.Int
}

func (frelconfig *Config) Read(input_file *file.File) *Config {

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	frelconfig.TimeSeconds = (&semantic.Int{}).Map(iniconfig.SectionMap[KEY_TIME_TAG][0], KEY_TIME_TAG, semantic.TypeVisible, inireader.REQUIRED_p)
	frelconfig.BasesMap = make(map[BaseNickname]*Base)
	frelconfig.Bases = make([]*Base, 0)
	frelconfig.SystemMap = make(map[SystemNickname]*System)
	frelconfig.Systems = make([]*System, 0)

	if bases, ok := iniconfig.SectionMap[KEY_BASE_TAG]; ok {
		for _, base := range bases {
			base_to_add := Base{}
			base_to_add.Map(base)
			base_to_add.Nickname = (&semantic.String{}).Map(base, KEY_NICKNAME, semantic.TypeVisible, inireader.REQUIRED_p)
			base_to_add.StridName = (&semantic.Int{}).Map(base, KEY_STRIDNAME, semantic.TypeVisible, inireader.REQUIRED_p)
			base_to_add.BGCS_base_run_by = (&semantic.String{}).Map(base, KEY_BASE_BGCS, semantic.TypeVisible, inireader.OPTIONAL_p)
			base_to_add.System = (&semantic.String{}).Map(base, KEY_SYSTEM, semantic.TypeVisible, inireader.REQUIRED_p)
			base_to_add.File = (&semantic.Path{}).Map(base, KEY_FILE, semantic.TypeVisible, inireader.REQUIRED_p)

			base_to_add.Nickname.Set(strings.ToLower(base_to_add.Nickname.Get()))
			base_to_add.System.Set(strings.ToLower(base_to_add.System.Get()))
			base_to_add.File.Set(strings.ToLower(base_to_add.File.Get()))

			base_to_add.Name = (&semantic.String{}).Map(base, KEY_NAME, semantic.TypeComment, inireader.OPTIONAL_p)
			base_to_add.RecycleCandidate = (&semantic.String{}).Map(base, KEY_RECYCLE, semantic.TypeComment, inireader.OPTIONAL_p)

			frelconfig.Bases = append(frelconfig.Bases, &base_to_add)
			frelconfig.BasesMap[BaseNickname(base_to_add.Nickname.Get())] = &base_to_add
		}
	}

	if systems, ok := iniconfig.SectionMap[KEY_SYSTEM_TAG]; ok {
		for _, system := range systems {
			system_to_add := System{}
			system_to_add.Map(system)

			system_to_add.Visit = (&semantic.Int{}).Map(system, KEY_SYSTEM_VISIT, semantic.TypeVisible, inireader.OPTIONAL_p)
			system_to_add.Strid_name = (&semantic.Int{}).Map(system, KEY_STRIDNAME, semantic.TypeVisible, inireader.OPTIONAL_p)
			system_to_add.Ids_info = (&semantic.Int{}).Map(system, KEY_SYSTEM_IDS_INFO, semantic.TypeVisible, inireader.OPTIONAL_p)
			// system_to_add.NavMapScale = system.GetParamNumber(KEY_SYSTEM_NAVMAPSCALE, inireader.OPTIONAL_p)
			system_to_add.Nickname = (&semantic.String{}).Map(system, KEY_NICKNAME, semantic.TypeVisible, inireader.REQUIRED_p)
			system_to_add.File = (&semantic.Path{}).Map(system, KEY_FILE, semantic.TypeVisible, inireader.OPTIONAL_p)
			system_to_add.Msg_id_prefix = (&semantic.String{}).Map(system, KEY_SYSTEM_MSG_ID_PREFIX, semantic.TypeVisible, inireader.OPTIONAL_p)

			system_to_add.Nickname.Set(strings.ToLower(system_to_add.Nickname.Get()))
			system_to_add.Msg_id_prefix.Set(strings.ToLower(system_to_add.Msg_id_prefix.Get()))

			if system_to_add.File.Get() != "" {
				system_to_add.File.Set(strings.ToLower(system_to_add.File.Get()))
			}

			frelconfig.Systems = append(frelconfig.Systems, &system_to_add)
			frelconfig.SystemMap[SystemNickname(system_to_add.Nickname.Get())] = &system_to_add
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

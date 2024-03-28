package systems_mapped

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/data_mapped/universe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

const (
	KEY_OBJECT   = "[Object]"
	KEY_NICKNAME = "nickname"
	KEY_BASE     = "base"
)

type Base struct {
	semantic.Model
	Nickname *semantic.String
	Base     *semantic.String // base.nickname in universe.ini
	DockWith *semantic.String

	IDsInfo     *semantic.Int
	IdsName     *semantic.Int
	RepNickname *semantic.String
}
type System struct {
	semantic.ConfigModel
	Nickname    string
	Bases       []*Base
	BasesByNick *lower_map.KeyLoweredMap[string, *Base]
}

type Config struct {
	SystemsMap *lower_map.KeyLoweredMap[string, *System]
	Systems    []*System
}

func Read(universe_config *universe_mapped.Config, filesystem *filefind.Filesystem) *Config {
	frelconfig := &Config{}

	var system_files map[string]*file.File = make(map[string]*file.File)
	for _, base := range universe_config.Bases {
		filename := universe_config.SystemMap.MapGet(universe_mapped.SystemNickname(base.System.Get())).File.FileName()
		path := filesystem.GetFile(utils_types.FilePath(strings.ToLower(filename)))
		system_files[base.System.Get()] = file.NewFile(path.GetFilepath())
	}

	var system_iniconfigs map[string]*inireader.INIFile = make(map[string]*inireader.INIFile)
	for system_key, file := range system_files {
		system_iniconfigs[system_key] = inireader.Read(file)
	}

	frelconfig.SystemsMap = lower_map.NewKeyLoweredMap[string, *System]()
	frelconfig.Systems = make([]*System, 0)
	for system_key, sysiniconf := range system_iniconfigs {
		system_to_add := &System{}
		system_to_add.Init(sysiniconf.Sections, sysiniconf.Comments, sysiniconf.File.GetFilepath())

		system_to_add.Nickname = system_key
		system_to_add.BasesByNick = lower_map.NewKeyLoweredMap[string, *Base]()
		system_to_add.Bases = make([]*Base, 0)
		frelconfig.SystemsMap.MapSet(system_key, system_to_add)
		frelconfig.Systems = append(frelconfig.Systems, system_to_add)

		if objects, ok := sysiniconf.SectionMap[KEY_OBJECT]; ok {
			for _, obj := range objects {

				// check if it is base object
				_, ok := obj.ParamMap[KEY_BASE]
				if ok {
					base_to_add := &Base{}
					base_to_add.Map(obj)

					base_to_add.Nickname = semantic.NewString(obj, KEY_NICKNAME)
					base_to_add.Base = semantic.NewString(obj, KEY_BASE)
					base_to_add.DockWith = semantic.NewString(obj, "dock_with", semantic.OptsS(semantic.Optional()))
					base_to_add.RepNickname = semantic.NewString(obj, "reputation", semantic.OptsS(semantic.Optional()))

					base_to_add.IDsInfo = semantic.NewInt(obj, "ids_info", semantic.Optional())
					base_to_add.IdsName = semantic.NewInt(obj, "ids_name", semantic.Optional())

					system_to_add.BasesByNick.MapSet(base_to_add.Nickname.Get(), base_to_add)
					system_to_add.Bases = append(system_to_add.Bases, base_to_add)
				}
			}
		}

	}

	return frelconfig
}

func (frelconfig *Config) Write() []*file.File {
	var files []*file.File = make([]*file.File, 0)
	for _, system := range frelconfig.Systems {
		inifile := system.Render()
		files = append(files, inifile.Write(inifile.File))
	}
	return files
}

package exe_mapped

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader/inireader_types"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
)

var KEY_BASE_TERRAINS = [...]string{"terrain_tiny", "terrain_sml", "terrain_mdm", "terrain_lrg", "terrain_dyna_01", "terrain_dyna_02"}

const (
	FILENAME_FL_INI                            = "freelancer.ini"
	RESOURCE_HEADER  inireader_types.IniHeader = "[Resources]"
	RESOURCE_KEY_DLL                           = "dll"
)

type Resources struct {
	semantic.Model
	Dll []string
}

type Config struct {
	semantic.ConfigModel

	Resources Resources
}

func (frelconfig *Config) Read(input_file *file.File) *Config {

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	if resources, ok := iniconfig.SectionMap[RESOURCE_HEADER]; ok {

		for _, dll := range resources[0].Params {
			frelconfig.Resources.Dll = append(frelconfig.Resources.Dll,
				strings.ReplaceAll(dll.First.AsString(), " ", ""),
			)
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

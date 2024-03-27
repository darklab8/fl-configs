package exe_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/go-utils/goutils/utils"
)

var KEY_BASE_TERRAINS = [...]string{"terrain_tiny", "terrain_sml", "terrain_mdm", "terrain_lrg", "terrain_dyna_01", "terrain_dyna_02"}

const (
	FILENAME_FL_INI = "freelancer.ini"
)

type Resources struct {
	semantic.Model
	Dll []*semantic.String
}

func (r *Resources) GetDlls() []string {
	return utils.CompL(r.Dll, func(x *semantic.String) string { return x.Get() })
}

type Config struct {
	semantic.ConfigModel

	Resources Resources
}

func Read(input_file *file.File) *Config {
	frelconfig := &Config{}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	if resources, ok := iniconfig.SectionMap["[Resources]"]; ok {

		for dll_index, _ := range resources[0].Params {
			frelconfig.Resources.Dll = append(frelconfig.Resources.Dll,
				semantic.NewString(resources[0], "dll", semantic.WithoutSpaces(), semantic.SOpts(semantic.Index(dll_index))),
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

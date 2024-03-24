package interface_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader/inireader_types"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
)

const (
	FILENAME_FL_INI                                     = "infocardmap.ini"
	RESOURCE_HEADER_MAP_TABLE inireader_types.IniHeader = "[InfocardMapTable]"
	RESOURCE_KEY_MAP                                    = "map"
)

type InfocardMapTable struct {
	semantic.Model
	Map map[string]string
}

type Config struct {
	semantic.ConfigModel

	InfocardMapTable InfocardMapTable
}

func Read(input_file *file.File) *Config {
	frelconfig := &Config{
		InfocardMapTable: InfocardMapTable{Map: make(map[string]string)},
	}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	if resources, ok := iniconfig.SectionMap[RESOURCE_HEADER_MAP_TABLE]; ok {

		for _, mappy := range resources[0].Params {

			frelconfig.InfocardMapTable.Map[mappy.First.AsString()] = mappy.Values[1].AsString()
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

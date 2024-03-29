package configfile

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
)

type ConfigFile struct {
	semantic.ConfigModel
	input_file *file.File
	Iniconfig  *inireader.INIFile
}

func NewConfigFile(input_file *file.File) *ConfigFile {
	fileconfig := &ConfigFile{input_file: input_file}
	return fileconfig
}

// Scan is heavy operations for goroutine ^_^
func (fileconfig *ConfigFile) Scan() *ConfigFile {
	iniconfig := inireader.Read(fileconfig.input_file)
	fileconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
	fileconfig.Iniconfig = iniconfig
	return fileconfig
}

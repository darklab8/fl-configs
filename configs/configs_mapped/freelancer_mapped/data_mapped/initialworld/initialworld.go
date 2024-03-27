package initialworld

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"
)

const (
	FILENAME = "initialworld.ini"
)

type Relationship struct {
	semantic.Model

	Rep            *semantic.Float
	TargetNickname *semantic.String
}

type Group struct {
	semantic.Model

	Nickname *semantic.String
	IdsName  *semantic.Int
	IdsInfo  *semantic.Int

	IdsShortName  *semantic.Int
	Relationships []*Relationship
}

type Config struct {
	semantic.ConfigModel

	Groups    []*Group
	GroupsMap *lower_map.KeyLoweredMap[string, *Group]
}

func Read(input_file *file.File) *Config {
	frelconfig := &Config{
		Groups:    make([]*Group, 0, 100),
		GroupsMap: lower_map.NewKeyLoweredMap[string, *Group](),
	}

	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())

	if groups, ok := iniconfig.SectionMap["[Group]"]; ok {

		for _, group_res := range groups {
			group := &Group{}
			group.Map(group_res)
			group.Nickname = semantic.NewString(group_res, "nickname")
			group.IdsName = semantic.NewInt(group_res, "ids_name")
			group.IdsInfo = semantic.NewInt(group_res, "ids_info")
			group.IdsShortName = semantic.NewInt(group_res, "ids_short_name")

			group.Relationships = make([]*Relationship, 0, 20)

			param_rep_key := "rep"
			for rep_index, _ := range group_res.ParamMap[param_rep_key] {

				rep := &Relationship{}
				rep.Map(group_res)
				rep.Rep = semantic.NewFloat(group_res, param_rep_key, semantic.Precision(2), semantic.Index(rep_index))
				rep.TargetNickname = semantic.NewString(group_res, param_rep_key, semantic.OptsS(semantic.Index(rep_index), semantic.Order(1)))
				group.Relationships = append(group.Relationships, rep)
			}

			frelconfig.Groups = append(frelconfig.Groups, group)
			frelconfig.GroupsMap.MapSet(group.Nickname.Get(), group)
		}
	}

	return frelconfig
}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

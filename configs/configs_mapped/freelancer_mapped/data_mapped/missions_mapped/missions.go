package missions_mapped

import (
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/semantic"
	"github.com/darklab8/fl-configs/configs/lower_map"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type EmpathyRate struct {
	semantic.Model

	TargetFactionNickname *semantic.String // 0
	RepoChange            *semantic.Float  // 1
}

type RepChangeEffects struct {
	semantic.Model
	Group *semantic.String

	EmpathyRates    []*EmpathyRate
	EmpathyRatesMap *lower_map.KeyLoweredMap[string, *EmpathyRate]
}

type Config struct {
	semantic.ConfigModel

	RepChangeEffects []*RepChangeEffects
	RepoChangeMap    *lower_map.KeyLoweredMap[string, *RepChangeEffects]
}

const (
	FILENAME utils_types.FilePath = "empathy.ini"
)

func Read(input_file *file.File) *Config {
	frelconfig := &Config{}
	iniconfig := inireader.INIFile.Read(inireader.INIFile{}, input_file)
	frelconfig.Init(iniconfig.Sections, iniconfig.Comments, iniconfig.File.GetFilepath())
	frelconfig.RepChangeEffects = make([]*RepChangeEffects, 0, 20)
	frelconfig.RepoChangeMap = lower_map.NewKeyLoweredMap[string, *RepChangeEffects]()

	for _, section := range iniconfig.SectionMap["[RepChangeEffects]"] {
		repo_changes := &RepChangeEffects{}
		repo_changes.Map(section)
		repo_changes.Group = semantic.NewString(section, "group")
		repo_changes.EmpathyRatesMap = lower_map.NewKeyLoweredMap[string, *EmpathyRate]()

		empathy_rate_key := "empathy_rate"
		for good_index, _ := range section.ParamMap[empathy_rate_key] {
			empathy := &EmpathyRate{}
			empathy.Map(section)
			empathy.TargetFactionNickname = semantic.NewString(section, empathy_rate_key, semantic.Index(good_index), semantic.Order(0))
			empathy.RepoChange = semantic.NewFloat(section, empathy_rate_key, semantic.Precision(2), semantic.Index(good_index), semantic.Order(1))
			repo_changes.EmpathyRates = append(repo_changes.EmpathyRates, empathy)
			repo_changes.EmpathyRatesMap.MapSet(empathy.TargetFactionNickname.Get(), empathy)
		}

		frelconfig.RepChangeEffects = append(frelconfig.RepChangeEffects, repo_changes)
		frelconfig.RepoChangeMap.MapSet(repo_changes.Group.Get(), repo_changes)
	}
	return frelconfig

}

func (frelconfig *Config) Write() *file.File {
	inifile := frelconfig.Render()
	inifile.Write(inifile.File)
	return inifile.File
}

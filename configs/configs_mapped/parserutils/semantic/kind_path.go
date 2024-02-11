package semantic

import (
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"
)

// Linux friendly filepath, that can be returned to Windows way from linux
type Path struct {
	Value
}

func NewPath(section *inireader.Section, key string, value_type ValueType, optional bool) *Path {
	s := &Path{}
	s.section = section
	s.key = key
	s.optional = optional
	s.value_type = value_type
	return s
}

func (s *Path) FileName() string {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return ""
	}
	result := s.section.ParamMap[s.key][0].First.AsString()
	return strings.Split(string(result), `\`)[2]
}

func (s *Path) Get() string {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return ""
	}
	return s.section.ParamMap[s.key][0].First.AsString()
}

func (s *Path) Set(value string) {
	if s.isComment() {
		s.Delete()
	}

	processed_value := inireader.UniParseStr(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment()}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *Path) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
		}
	}
}

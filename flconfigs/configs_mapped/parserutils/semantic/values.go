/*
ORM mapper for Freelancer ini reader. Easy mapping values to change.
*/
package semantic

import (
	"strings"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/inireader"
)

// ORM values

const (
	TypeComment = true
	TypeVisible = false
)

type Value struct {
	section   *inireader.Section
	key       string
	optional  bool
	isComment bool
}

type String struct {
	Value
}

func (s *String) Map(section *inireader.Section, key string, isComment bool, optional bool) *String {
	s.section = section
	s.key = key
	s.optional = optional
	s.isComment = isComment
	return s
}

func (s *String) Get() string {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return ""
	}
	return s.section.ParamMap[s.key][0].First.AsString()
}

func (s *String) Set(value string) {
	if s.isComment {
		s.Delete()
	}

	processed_value := inireader.UniParseStr(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *String) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
		}
	}
}

type Int struct {
	Value
}

func (s *Int) Map(section *inireader.Section, key string, isComment bool, optional bool) *Int {
	s.section = section
	s.key = key
	s.optional = optional
	s.isComment = isComment
	return s
}

func (s *Int) Get() int {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return 0
	}
	return int(s.section.ParamMap[s.key][0].First.(inireader.ValueNumber).Value)
}

func (s *Int) Set(value int) {
	if s.isComment {
		s.Delete()
	}

	processed_value := inireader.UniParseInt(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *Int) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
		}
	}
}

// Linux friendly filepath, that can be returned to Windows way from linux
type Path struct {
	Value
}

func (s *Path) Map(section *inireader.Section, key string, isComment bool, optional bool) *Path {
	s.section = section
	s.key = key
	s.optional = optional
	s.isComment = isComment
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
	if s.isComment {
		s.Delete()
	}

	processed_value := inireader.UniParseStr(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment}).AddValue(processed_value))
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

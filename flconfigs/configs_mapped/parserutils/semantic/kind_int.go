package semantic

import "github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/inireader"

type Int struct {
	Value
}

func NewInt(section *inireader.Section, key string, value_type ValueType, optional bool) *Int {
	s := &Int{}
	s.section = section
	s.key = key
	s.optional = optional
	s.value_type = value_type
	return s
}

func (s *Int) Get() int {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return 0
	}
	return int(s.section.ParamMap[s.key][0].First.(inireader.ValueNumber).Value)
}

func (s *Int) Set(value int) {
	if s.isComment() {
		s.Delete()
	}

	processed_value := inireader.UniParseInt(value)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment()}).AddValue(processed_value))
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

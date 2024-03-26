package semantic

import "github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"

type Bool struct {
	*Value
}

func NewBool(section *inireader.Section, key string, opts ...ValueOption) *Bool {
	v := NewValue(section, key)
	for _, opt := range opts {
		opt(v)
	}
	s := &Bool{Value: v}

	return s
}

func (s *Bool) Get() bool {
	if s.optional && len(s.section.ParamMap[s.key]) == 0 {
		return false
	}
	return int(s.section.ParamMap[s.key][s.index].Values[s.order].(inireader.ValueNumber).Value) == 1
}

func (s *Bool) Set(value bool) {
	var int_bool int

	if value {
		int_bool = 1
	}

	if s.isComment() {
		s.Delete()
	}

	processed_value := inireader.UniParseInt(int_bool)
	if len(s.section.ParamMap[s.key]) == 0 {
		s.section.AddParamToStart(s.key, (&inireader.Param{IsComment: s.isComment()}).AddValue(processed_value))
	}
	// implement SetValue in Section
	s.section.ParamMap[s.key][0].First = processed_value
	s.section.ParamMap[s.key][0].Values[0] = processed_value
}

func (s *Bool) Delete() {
	delete(s.section.ParamMap, s.key)
	for index, param := range s.section.Params {
		if param.Key == s.key {
			s.section.Params = append(s.section.Params[:index], s.section.Params[index+1:]...)
		}
	}
}

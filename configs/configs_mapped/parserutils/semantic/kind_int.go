package semantic

import "github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/inireader"

type Int struct {
	*Value
}

type IntOption func(i *Int)

func IntOpts(opts ...ValueOption) IntOption {
	return func(i *Int) {
		for _, opt := range opts {
			opt(i.Value)
		}
	}
}

func NewInt(section *inireader.Section, key string, opts ...IntOption) *Int {
	s := &Int{Value: NewValue(section, key)}
	for _, opt := range opts {
		opt(s)
	}

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

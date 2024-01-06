/*
ORM mapper for Freelancer ini reader. Easy mapping values to change.
*/
package semantic

import (
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/inireader"
)

// ORM values

type ValueType int64

const (
	TypeComment ValueType = iota
	TypeVisible
)

type Value struct {
	section    *inireader.Section
	key        string
	optional   bool
	value_type ValueType
}

func (v Value) isComment() bool {
	return v.value_type == TypeComment
}

package inireader

import (
	"regexp"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils"
)

func initRegexExpression(regex **regexp.Regexp, expression string) {
	var err error

	*regex, err = regexp.Compile(expression)
	logus.Log.CheckFatal(err, "failed to parse numberParser in ", logus_core.FilePath(utils.GetCurrentFile()))
}

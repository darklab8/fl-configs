package inireader

import (
	"regexp"

	"github.com/darklab8/darklab_flconfigs/flconfigs/settings/logger"

	"github.com/darklab8/darklab_goutils/goutils/utils"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_logger"
)

func initRegexExpression(regex **regexp.Regexp, expression string) {
	var err error

	*regex, err = regexp.Compile(expression)
	logger.Log.CheckFatal(err, "failed to parse numberParser in ", utils_logger.FilePath(utils.GetCurrentFile()))
}

package inireader

import (
	"regexp"

	"github.com/darklab8/fl-configs/configs/settings/logger"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
)

func initRegexExpression(regex **regexp.Regexp, expression string) {
	var err error

	*regex, err = regexp.Compile(expression)
	logger.Log.CheckFatal(err, "failed to parse numberParser in ", utils_logus.FilePath(utils.GetCurrentFile()))
}

/*
Such code is primiarily used for fl-darklint. You could check its code for more examples
https://github.com/darklab8/fl-darklint
*/
package configs

import (
	"os"
	"strings"

	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

var Freelancerfolder utils_types.FilePath = utils_types.FilePath(os.Getenv("CONFIGS_FREELANCER_FOLDER"))

// ExampleModifyingData demononstrating how to change configs values
func Example_modifyingConfigs() {
	configs := configs_mapped.NewMappedConfigs()
	logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(Freelancerfolder))

	// Reading ini reading universal format
	// and mapping to ORM objects
	configs.Read(FreelancerFolder)

	// Modifying files
	for _, base := range configs.Universe_config.Bases {
		base.Nickname.Set(strings.ToLower(base.Nickname.Get()))
		base.System.Set(strings.ToLower(base.System.Get()))
		base.File.Set(strings.ToLower(base.File.Get()))
	}

	for _, system := range configs.Universe_config.Systems {
		system.Nickname.Set(strings.ToLower(system.Nickname.Get()))
		system.Msg_id_prefix.Set(strings.ToLower(system.Msg_id_prefix.Get()))

		if system.File.Get() != "" {
			system.File.Set(strings.ToLower(system.File.Get()))
		}
	}

	// Write without Dry Run for writing to files modified values back!
	configs.Write(configs_mapped.IsDruRun(true))
}

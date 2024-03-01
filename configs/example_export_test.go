/*
Such code is primiarily used for fl-darkstat. You could check its code for more examples
https://github.com/darklab8/fl-darkstat
*/
package configs

import (
	"fmt"
	"os"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

var FreelancerFolder utils_types.FilePath = utils_types.FilePath(os.Getenv("CONFIGS_FREELANCER_FOLDER"))

// ExampleExportingData demononstrating exporting freelancer folder data for comfortable usage
func Example_exportingData() {
	configs := configs_mapped.NewMappedConfigs()
	logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(FreelancerFolder))

	// Reading to ini universal custom format and mapping to ORM objects
	// which have both reading and writing back capabilities
	configs.Read(FreelancerFolder)

	// For elegantly exporting enriched data objects with better type safety for just reading access
	exports := configs_export.NewExporter(configs)

	bases := exports.Bases(configs_export.NoNameIncluded(false))
	for _, base := range bases {
		// do smth with exported bases
		fmt.Println(base.Name)
		fmt.Println(base.Infocard)
		fmt.Println(base.System)
		fmt.Println(base.SystemNickname)
		fmt.Printf("%d\n", base.InfocardID)
		break
	}
}

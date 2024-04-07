/*
See package `configs` for description and code examples
*/
package main

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/darklab8/fl-configs/configs/configs_export"
	"github.com/darklab8/fl-configs/configs/configs_mapped"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func main() {

	// for profiling only stuff.
	f, err := os.Create("prof.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 1; i++ {
		time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
			var configs *configs_mapped.MappedConfigs
			time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
				real_game_loc := utils_types.FilePath(os.Getenv("FREELANCER_FOLDER"))

				configs = configs_mapped.NewMappedConfigs()
				logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(real_game_loc))
				configs.Read(real_game_loc)
			}, time_measure.WithMsg("read mapping"))
			time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
				exported := configs_export.Export(configs)

				// config := exe_mapped.FixtureFLINIConfig()
				// ids := exe_mapped.GetAllInfocards(filefind.FindConfigs(real_game_loc), config.GetDlls())

				for range exported.Bases.Bases {
					break
				}
			}, time_measure.WithMsg("exported time"))
		}, time_measure.WithMsg("total time"))
	}
}

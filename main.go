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
	"github.com/darklab8/fl-configs/configs/configs_settings"
	"github.com/darklab8/fl-configs/configs/configs_settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
	"github.com/darklab8/go-utils/goutils/utils/utils_logus"
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
				freelancer_folder := configs_settings.GetGameLocation()

				configs = configs_mapped.NewMappedConfigs()
				logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(freelancer_folder))
				configs.Read(freelancer_folder)
			}, time_measure.WithMsg("read mapping"))
			time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
				exported := configs_export.Export(configs)

				// config := exe_mapped.FixtureFLINIConfig()
				// ids := exe_mapped.GetAllInfocards(filefind.FindConfigs(real_game_loc), config.GetDlls())

				for range exported.Bases {
					break
				}
			}, time_measure.WithMsg("exported time"))
		}, time_measure.WithMsg("total time"))
	}
}

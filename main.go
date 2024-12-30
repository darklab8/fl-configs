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
	"github.com/darklab8/go-utils/utils/timeit"
	"github.com/darklab8/go-utils/utils/utils_logus"
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
		timeit.NewTimerF(func() {
			var configs *configs_mapped.MappedConfigs
			timeit.NewTimerF(func() {
				freelancer_folder := configs_settings.Env.FreelancerFolder

				configs = configs_mapped.NewMappedConfigs()
				logus.Log.Debug("scanning freelancer folder", utils_logus.FilePath(freelancer_folder))
				configs.Read(freelancer_folder)
			}, timeit.WithMsg("read mapping"))
			timeit.NewTimerF(func() {
				exported := configs_export.Export(configs, configs_export.ExportOptions{})

				// config := exe_mapped.FixtureFLINIConfig()
				// ids := exe_mapped.GetAllInfocards(filefind.FindConfigs(real_game_loc), config.GetDlls())

				for range exported.Bases {
					break
				}
			}, timeit.WithMsg("exported time"))
		}, timeit.WithMsg("total time"))
	}
}

/*
See package `configs` for description and code examples
*/
package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"

	"github.com/darklab8/fl-configs/configs/configs_mapped/configs_fixtures"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
)

func main() {

	// for profiling only stuff.
	f, err := os.Create("prof.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
		game_location := configs_fixtures.FixtureGameLocation()
		config := exe_mapped.FixtureFLINIConfig()
		ids := exe_mapped.GetAllInfocards(filefind.FindConfigs(game_location), config.GetDlls())

		for id, text := range ids.Infocards {
			fmt.Println(id)
			fmt.Println(text)
			break
		}
	})
}

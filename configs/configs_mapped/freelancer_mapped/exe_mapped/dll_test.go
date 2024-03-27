package exe_mapped

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/configs_fixtures"
	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/infocard_mapped/infocard"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
	"github.com/stretchr/testify/assert"
)

func TestReadInfocards(t *testing.T) {
	game_location := configs_fixtures.FixtureGameLocation()
	config := FixtureFLINIConfig()
	dlls := config.Resources.GetDlls()
	infocards := GetAllInfocards(filefind.FindConfigs(game_location), dlls)

	assert.Greater(t, len(infocards.Infocards), 0)
	assert.Greater(t, len(infocards.Infonames), 0)

	for id, text := range infocards.Infonames {
		fmt.Println(id)
		fmt.Println(text)
		break
	}

	assert.Contains(t, infocards.Infocards[132903].Content, "We just brought a load of Fertilizers")

	fmt.Println(infocards.Infocards[196624])
	fmt.Println("second:", infocards.Infocards[66089])

	fmt.Println("Abandoned Depot infocard\n",
		infocards.Infocards[465639],
		infocards.Infocards[465639+1], // value from infocardmap.txt mapped
		infocards.Infocards[500904],   // faction infocard id
	)
}

func TestReadInfocardsToHtml(t *testing.T) {
	f, err := os.Create("prof.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	result := time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
		game_location := configs_fixtures.FixtureGameLocation()
		config := FixtureFLINIConfig()
		infocards := GetAllInfocards(filefind.FindConfigs(game_location), config.Resources.GetDlls())

		// assert.Greater(t, len(ids), 0)

		// 503718 faction BMM
		// 465639 base Bandoned Depot
		// 465640 continuation
		// infocard tail 500904

		xml_stuff := infocards.Infocards[501545]
		fmt.Println("xml_stuff=", xml_stuff)

		text, err := xml_stuff.XmlToText()
		logus.Log.CheckFatal(err, "unable convert to text")

		assert.Greater(t, len(text), 0)
		assert.NotEmpty(t, text)
		fmt.Println(text)

	}, time_measure.WithMsg("measure time"))
	logus.Log.CheckFatal(result.ResultErr, "non nil exit")
}

func TestValidateInfocards(t *testing.T) {
	game_location := configs_fixtures.FixtureGameLocation()
	config := FixtureFLINIConfig()
	infocards := GetAllInfocards(filefind.FindConfigs(game_location), config.Resources.GetDlls())

	var parsed []*infocard.Infocard = make([]*infocard.Infocard, 0, 100)
	var parsed_text map[int][]string = make(map[int][]string)
	var failed []*infocard.Infocard = make([]*infocard.Infocard, 0, 100)

	for id, infocard := range infocards.Infocards {
		text, err := infocard.XmlToText()
		parsed_text[id] = text

		if logus.Log.CheckWarn(err, "unable convert to text") {
			failed = append(failed, infocard)
		} else {
			parsed = append(parsed, infocard)
		}
	}

	fmt.Println("parsed_count=", len(parsed))
	assert.Equal(t, len(failed), 0, "expected no failed")
}

func TestCoversion(t *testing.T) {
	windows_decoded := "A\x00L\x00L\x00I\x00E\x00S\x00:\x00\u00a0\x00 \x00<\x00/\x00T\x00E\x00X\x00T\x00>\x00<\x00P\x00A\x00R"

	sliced := make([]byte, 0, len(windows_decoded)/2)
	// sliced := make([]rune, len(str_windows_decoded)/2)
	for i := 0; i < len(windows_decoded); i += 2 {
		sliced = append(sliced, windows_decoded[i]) // or do whatever
	}

	fmt.Println(string(sliced))

	decoded, _ := DecodeUTF16([]byte(windows_decoded))
	fmt.Println(decoded)
}

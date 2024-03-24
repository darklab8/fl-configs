package exe_mapped

import (
	"encoding/xml"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/darklab8/fl-configs/configs/configs_mapped/configs_fixtures"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-utils/goutils/utils/time_measure"
	"github.com/stretchr/testify/assert"
)

func TestReadInfocards(t *testing.T) {
	game_location := configs_fixtures.FixtureGameLocation()
	config := FixtureFLINIConfig()
	ids := GetAllInfocards(filefind.FindConfigs(game_location), config.Resources.Dll)

	assert.Greater(t, len(ids), 0)

	for id, text := range ids {
		fmt.Println(id)
		fmt.Println(text)
		break
	}

	assert.Contains(t, ids[132903], "We just brought a load of Fertilizers")

	fmt.Println(ids[196624])
	fmt.Println("second:", ids[66089])

	fmt.Println("Abandoned Depot infocard\n",
		ids[465639],
		ids[465639+1], // value from infocardmap.txt mapped
		ids[500904],   // faction infocard id
	)
}

type RDL struct {
	XMLName xml.Name `xml:"RDL"`
	Text    string   `xml:",chardata"`
	PUSH    string   `xml:"PUSH"`
	TEXT    []string `xml:"TEXT"`
	PARA    []string `xml:"PARA"`
	POP     string   `xml:"POP"`
}

type RDL2 struct {
	XMLName xml.Name `xml:"RDL"`
	TEXT    []string `xml:"TEXT"`
}

func TestReadInfocardsToHtml(t *testing.T) {
	result := time_measure.TimeMeasure(func(m *time_measure.TimeMeasurer) {
		game_location := configs_fixtures.FixtureGameLocation()
		config := FixtureFLINIConfig()
		ids := GetAllInfocards(filefind.FindConfigs(game_location), config.Resources.Dll)

		assert.Greater(t, len(ids), 0)

		// 503718 faction BMM
		// 465639 base Bandoned Depot
		// 465640 continuation
		// infocard tail 500904

		xml_stuff := ids[465639]
		fmt.Println("xml_stuff=", xml_stuff)
		var mappy map[string]interface{} = make(map[string]interface{})
		_ = mappy
		var structy RDL2
		_ = structy
		err := xml.Unmarshal([]byte(strings.ReplaceAll(string(xml_stuff), `<?xml version="1.0" encoding="UTF-16"?>`, "")), &structy)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(structy)

	}, time_measure.WithMsg("measure time"))
	logus.Log.CheckFatal(result.ResultErr, "non nil exit")
}

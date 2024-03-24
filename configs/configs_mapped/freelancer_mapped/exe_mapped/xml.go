package exe_mapped

import (
	"encoding/xml"
	"strings"
)

type RDL struct {
	XMLName xml.Name `xml:"RDL"`
	TEXT    []string `xml:"TEXT"`
}

func XmlToText(xml_stuff string) ([]string, error) {
	var structy RDL
	err := xml.Unmarshal([]byte(strings.ReplaceAll(string(xml_stuff), `<?xml version="1.0" encoding="UTF-16"?>`, "")), &structy)

	// logus.Log.CheckError(err, "unable converting xml to text", typelog.String("xml_stuff", string(xml_stuff)))
	if err != nil {
		return nil, err
	}
	return structy.TEXT, nil
}

package infocard

import (
	"bytes"
	"encoding/xml"
	"io"
)

func (i *Infocard) XmlToText() ([]string, error) {
	return XmlToText(i.content)
}

func XmlToText(raw string) ([]string, error) {
	decoder := xml.NewDecoder(bytes.NewBufferString(raw))

	lines := make([]string, 0)
	line := ""
	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		switch tok := token.(type) {
		case xml.EndElement:
			if tok.Name.Local == "PARA" || tok.Name.Local == "POP" {
				lines = append(lines, line)
				line = ""
			}
		case xml.CharData:
			line += string(tok)
		default:
			continue
		}
	}

	return lines, nil
}

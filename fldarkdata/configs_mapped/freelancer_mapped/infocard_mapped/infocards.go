package infocard_mapped

import (
	"strconv"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
)

type Record interface {
	Id() int
	Type() string
	Content() string
}
type Infocard struct {
	id      int
	content string
}
type Name struct {
	id      int
	content string
}

func (v Name) Id() int {
	return v.id
}
func (v Infocard) Id() int {
	return v.id
}
func (v Name) Content() string {
	return v.content
}
func (v Infocard) Content() string {
	return v.content
}

func (v Name) Type() string {
	return TYPE_NAME
}
func (v Infocard) Type() string {
	return TYPE_INFOCAD
}

type Config struct {
	Records []*Record

	RecordsMap map[int]*Record
}

const (
	FILENAME          = "infocards.txt"
	FILENAME_FALLBACK = "infocards.xml"
	TYPE_NAME         = "NAME"
	TYPE_INFOCAD      = "INFOCARD"
)

func (frelconfig *Config) Read(input_file *file.File) *Config {
	frelconfig.RecordsMap = make(map[int]*Record)
	frelconfig.Records = make([]*Record, 0)

	input_file = input_file.OpenToReadF()
	defer input_file.Close()
	lines := input_file.ReadLines()

	for index := 0; index < len(lines); index++ {

		id, err := strconv.Atoi(lines[index])
		if err != nil {
			continue
		}
		name := lines[index+1]
		content := lines[index+2]
		index += 3

		var record_to_add Record
		switch name {
		case TYPE_NAME:
			record_to_add = Name{id: id, content: content}
		case TYPE_INFOCAD:
			record_to_add = Infocard{id: id, content: content}
		default:
			logus.Log.Fatal(
				"unrecognized object name in infocards.txt",
				logus_core.Any("id", id),
				logus_core.Any("name", name),
				logus_core.Any("content", content),
			)
		}

		frelconfig.Records = append(frelconfig.Records, &record_to_add)
		frelconfig.RecordsMap[record_to_add.Id()] = &record_to_add
	}

	return frelconfig
}

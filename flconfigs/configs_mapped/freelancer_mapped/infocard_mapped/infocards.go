package infocard_mapped

import (
	"strconv"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_flconfigs/flconfigs/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
)

type Record struct {
	Id      int
	Content string
	Kind    RecordKind
}

func NewRecord(Id int, Content string, Kind RecordKind) *Record {
	return &Record{Id: Id, Content: Content, Kind: Kind}
}

type RecordKind string

const (
	TYPE_NAME    RecordKind = "NAME"
	TYPE_INFOCAD RecordKind = "INFOCARD"
)

type Config struct {
	Records []*Record

	RecordsMap map[int]*Record
}

const (
	FILENAME          = "infocards.txt"
	FILENAME_FALLBACK = "infocards.xml"
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

		var record_to_add *Record
		switch RecordKind(name) {
		case TYPE_NAME:
			record_to_add = NewRecord(id, content, TYPE_NAME)
		case TYPE_INFOCAD:
			record_to_add = NewRecord(id, content, TYPE_INFOCAD)
		default:
			logus.Log.Fatal(
				"unrecognized object name in infocards.txt",
				logus_core.Any("id", id),
				logus_core.Any("name", name),
				logus_core.Any("content", content),
			)
		}

		frelconfig.Records = append(frelconfig.Records, record_to_add)
		frelconfig.RecordsMap[record_to_add.Id] = record_to_add
	}

	return frelconfig
}

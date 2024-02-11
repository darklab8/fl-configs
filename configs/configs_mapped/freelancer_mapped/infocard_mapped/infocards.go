package infocard_mapped

import (
	"strconv"

	"github.com/darklab8/fl-configs/configs/configs_mapped/freelancer_mapped/exe_mapped"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/fl-configs/configs/configs_mapped/parserutils/filefind/file"
	logus1 "github.com/darklab8/fl-configs/configs/settings/logus"
	"github.com/darklab8/go-typelog/typelog"
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
	TYPE_NAME      RecordKind = "NAME"
	TYPE_INFOCAD   RecordKind = "INFOCARD"
	TYPE_UNDEFINED RecordKind = "not parsed yet"
)

type Config struct {
	Records    []*Record
	RecordsMap map[int]*Record
}

const (
	FILENAME          = "infocards.txt"
	FILENAME_FALLBACK = "infocards.xml"
)

func (frelconfig *Config) ReadFromTextFile(input_file *file.File) {

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
			logus1.Log.Fatal(
				"unrecognized object name in infocards.txt",
				typelog.Any("id", id),
				typelog.Any("name", name),
				typelog.Any("content", content),
			)
		}

		frelconfig.Records = append(frelconfig.Records, record_to_add)
		frelconfig.RecordsMap[record_to_add.Id] = record_to_add
	}
}

func (frelconfig *Config) Read(filesystem *filefind.Filesystem, freelancer_ini *exe_mapped.Config, input_file *file.File) *Config {
	frelconfig.RecordsMap = make(map[int]*Record)
	frelconfig.Records = make([]*Record, 0)

	// p.Infocards =
	if input_file != nil {
		frelconfig.ReadFromTextFile(input_file)
	} else {
		infocards := exe_mapped.GetAllInfocards(filesystem, freelancer_ini.Resources.Dll)
		for id, text := range infocards {
			var record_to_add *Record = NewRecord(int(id), string(text), TYPE_INFOCAD)
			frelconfig.Records = append(frelconfig.Records, record_to_add)
			frelconfig.RecordsMap[record_to_add.Id] = record_to_add
		}
	}

	return frelconfig
}

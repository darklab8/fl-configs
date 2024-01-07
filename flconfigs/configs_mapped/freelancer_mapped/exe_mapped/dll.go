package exe_mapped

import (
	"math"
	"os"
	"strings"

	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind"
	"github.com/darklab8/darklab_flconfigs/flconfigs/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_flconfigs/flconfigs/settings/logus"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type InfocardID string
type InfocardText string

const SEEK_SET = 0 // python default seek(offset, whence=os.SEEK_SET, /)¶

func parseDLL(fh *os.File, out map[InfocardID]InfocardText, global_offset int) {
	fh.Seek(0x3C, SEEK_SET)
	// var Array1Len []byte = make([]byte, 1)
	PE_sig_loc, err := fh.Read(make([]byte, 1))
	logus.Log.CheckError(err, "failed reading PE_sig_loc")
	fh.Seek(int64(PE_sig_loc+4), SEEK_SET)
}

func ParseDLLs(dll_fnames []*file.File) map[InfocardID]InfocardText {
	out := make(map[InfocardID]InfocardText, 0)

	for idx, name := range dll_fnames {
		fh, err := os.Open(name.GetFilepath().ToString())

		if logus.Log.CheckError(err, "unable to read dll") {
			continue
		}

		global_offset := int(math.Pow(2, 16)) * (idx + 1)
		parseDLL(fh, out, global_offset)
	}

	return out
}

func GetAllInfocards(game_location utils_types.FilePath, dll_names []string) map[InfocardID]InfocardText {

	var files []*file.File
	filesystem := filefind.FindConfigs(game_location)
	for _, filename := range dll_names {
		dll_file := filesystem.GetFile(utils_types.FilePath(strings.ToLower(filename)))
		files = append(files, dll_file)
	}

	return ParseDLLs(files)
}

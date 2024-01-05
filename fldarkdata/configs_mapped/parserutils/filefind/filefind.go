/*
Package with reusable code for discovery of files and other reusable stuff like universal ini reader
*/
package filefind

import (
	"io/fs"
	"path/filepath"
	"strings"

	"github.com/darklab8/darklab_fldarkdata/fldarkdata/configs_mapped/parserutils/filefind/file"
	"github.com/darklab8/darklab_fldarkdata/fldarkdata/settings/logus"

	"github.com/darklab8/darklab_goutils/goutils/logus_core"
	"github.com/darklab8/darklab_goutils/goutils/utils/utils_types"
)

type Filesystem struct {
	Files   []*file.File
	Hashmap map[utils_types.FilePath]*file.File
}

var FreelancerFolder Filesystem

func FindConfigs(folderpath utils_types.FilePath) Filesystem {
	var filesystem Filesystem
	filesystem.Hashmap = make(map[utils_types.FilePath]*file.File)

	err := filepath.WalkDir(string(folderpath), func(path string, d fs.DirEntry, err error) error {

		if !strings.Contains(path, ".ini") && !strings.Contains(path, ".txt") && !strings.Contains(path, ".xml") {
			return nil
		}

		logus.Log.CheckFatal(err, "unable to read file")

		file := file.NewFile(utils_types.FilePath(path))
		filesystem.Files = append(filesystem.Files, file)

		key := utils_types.FilePath(strings.ToLower(filepath.Base(path)))
		filesystem.Hashmap[key] = file

		return nil
	})

	logus.Log.CheckFatal(err, "unable to read files")
	return filesystem
}

func (file1system Filesystem) GetFile(file1names ...utils_types.FilePath) *file.File {
	for _, file1name := range file1names {
		file_, ok := file1system.Hashmap[file1name]
		if !ok {
			logus.Log.Warn("Filesystem.GetFile, failed to find find in filesystesm file trying to recover", logus_core.FilePath(file1name))
			continue
		}
		logus.Log.Info("Filesystem.GetFile, found filepath=", logus_core.FilePath(file_.GetFilepath()))
		result_file := file.NewFile(file_.GetFilepath())
		return result_file
	}
	logus.Log.Fatal("unable to find filenames=", logus_core.Filepaths(file1names))
	return nil
}

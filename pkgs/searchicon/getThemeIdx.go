package searchicon

import (
	"fmt"
	"io/fs"
	"os"
	"path"
)

// Returns the first index.theme file found in theme directories
func getThemeIdx(themeDirs string) (*os.File, error) {
	var idxFile *os.File
	var err error = fmt.Errorf("index.theme not found in %v", themeDirs)

	walkDirsStr(themeDirs, ':', func(dir string) error {
		iIdxFile, iErr := os.Open(path.Join(dir, "index.theme"))
		if iErr == nil {
			idxFile = iIdxFile
			err = nil
			return fs.SkipAll
		}
		return nil
	})

	return idxFile, err
}

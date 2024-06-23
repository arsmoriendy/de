package searchicon

import (
	"errors"
	"os"
	"path"
)

// Returns colon separated dir strings
//
// # Returns empty string if non found
//
// WARNING: make sure baseDirs is initialized
func getThemeDirs(themeName string) string {
	rstring := ""

	walkDirsStr(baseDirs, ':', func(baseDir string) error {
		themeDir := path.Join(baseDir, themeName)

		if _, err := os.Stat(themeDir); errors.Is(err, os.ErrNotExist) {
			return nil
		}

		if rstring != "" {
			rstring += ":"
		}

		rstring += themeDir

		return nil
	})

	return rstring
}
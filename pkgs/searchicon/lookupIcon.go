package searchicon

import (
	"io/fs"
	"os"
	"path"
)

// iconname does not include the file extension (e.g. "icon" not "icon.svg")
func lookupIcon(iconname string, size int, scale int, theme string) (string, error) {
	rstring := ""

	idxFile, err := getThemeIdx(getThemeDirs(theme))
	if err != nil {
		return "", err
	}

	subdirs, err := GetHKV(idxFile.Name(), "Icon Theme", "Directories")
	if err != nil {
		return "", err
	}

	// Append directories listed in ScaledDirectories if any for backwards compatibility
	scaledSubdirs, err := GetHKV(idxFile.Name(), "Icon Theme", "ScaledDirectories")
	if err == nil {
		subdirs += "," + scaledSubdirs
	}

	walkDirsStr(subdirs, ',', func(subdir string) error {
		var wds1rerr error = nil

		matches, _ := dirMatchesSize(idxFile, subdir, size, scale)
		if !matches {
			return nil
		}

		walkDirsStr(baseDirs, ':', func(directory string) error {
			dirpath := path.Join(directory, theme, subdir)

			for _, ext := range []string{"png", "svg", "xpm"} {
				filename := path.Join(dirpath, iconname+"."+ext)

				if _, err = os.Stat(filename); err != nil {
					continue
				}

				rstring = filename
				wds1rerr = fs.SkipAll
				return fs.SkipAll
			}

			return nil
		})

		return wds1rerr
	})

	// TODO: DirectorySizeDistance loop

	return rstring, nil
}

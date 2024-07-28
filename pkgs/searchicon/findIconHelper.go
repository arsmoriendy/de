package searchicon

import (
	"io/fs"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

func findIconHelper(icon string, size int, scale int, theme string) (string, error) {
	filename, err := lookupIcon(icon, size, scale, theme)
	if err == nil {
		return filename, err
	}

	idxFile, err := getThemeIdx(getThemeDirs(theme))
	if err != nil {
		return "", IconNotFound
	}

	ispec := iconspec.New(idxFile.Name())

	parents, err := ispec.Inherits()
	if err != nil {
		return "", IconNotFound
	}

	walkDirsStr(parents, ',', func(theme string) error {
		if theme == "hicolor" {
			return nil
		}

		filename, err = findIconHelper(icon, size, scale, theme)
		if err == nil {
			return fs.SkipAll
		}

		return nil
	})

	if filename == "" {
		return filename, IconNotFound
	} else {
		return filename, nil
	}
}

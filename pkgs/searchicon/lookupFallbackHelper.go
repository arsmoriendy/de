package searchicon

import (
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

func lookupFallbackHelper(iconname string, theme string) (string, error) {
	icon, err := findFallback(iconname, theme)
	if icon != "" && err == nil {
		return icon, nil
	}

	icon, err = findFallback(iconname, "hicolor")
	if icon != "" && err == nil {
		return icon, nil
	}

	return icon, IconNotFound
}

func findFallback(iconname string, theme string) (string, error) {
	rstring := ""

	themedirs := getThemeDirs(theme)

	idx, err := getThemeIdx(themedirs)
	if err != nil {
		return rstring, fmt.Errorf("%w: %w", IconNotFound, err)
	}

	is := iconspec.New(idx.Name())

	dirs, err := is.Directories()
	if err != nil {
		return rstring, fmt.Errorf("%w: %w", IconNotFound, err)
	}

	if sdirs, err := is.ScaledDirectories(); sdirs != "" && err == nil {
		dirs += "," + sdirs
	}

	walkDirsStr(dirs, ',', func(dir string) error {
		skipall := false

		walkDirsStr(themedirs, ':', func(themedir string) error {
			// TODO: test this
			path := path.Join(themedir, dir)
			icon, err := lookupFallbackIcon(iconname, path)
			if err != nil {
				return nil
			}

			if _, err := os.Stat(icon); err != nil {
				return nil
			}

			rstring = icon
			skipall = true
			return fs.SkipAll
		})

		if skipall {
			return fs.SkipAll
		}

		return nil
	})

	if rstring != "" {
		return rstring, nil
	}

	return rstring, IconNotFound
}

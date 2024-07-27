package searchicon

import (
	"io/fs"
	"os"
	"path"
)

// Lookup icon regardless of size or scale
func lookupFallbackIcon(iconname string, basenames string) (string, error) {
	ricon := ""
	walkDirsStr(basenames, ',', func(dir string) error {
		for _, ext := range []string{"png", "svg", "xpm"} {
			fn := path.Join(dir, iconname+"."+ext)

			if _, err := os.Stat(fn); err != nil {
				continue
			}

			ricon = fn
			return fs.SkipAll
		}
		return nil
	})

	if ricon == "" {
		return ricon, IconNotFound
	} else {
		return ricon, nil
	}
}

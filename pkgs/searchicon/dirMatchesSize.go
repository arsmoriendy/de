package searchicon

import (
	"errors"
	"os"

	"github.com/arsmoriendy/de/pkgs/searchicon/gethkv"
	"github.com/arsmoriendy/de/pkgs/searchicon/iconspec"
)

var dirSizeMismatch = errors.New("theme subdirectory does not match the size constraint specified.")
var idxFormatErr = errors.New("index.theme file has an invalid format")

// Where idxFile is the index.theme file.
//
// subdir is not an absolute path, it is one of the headers in index.theme
func dirMatchesSize(idxFile *os.File, subdir string, iconsize int, iconscale int) (bool, error) {
	ds := iconspec.NewDir(idxFile.Name(), subdir)

	// if Scale != iconscale [
	scale := ds.Scale()

	if scale != iconscale {
		return false, nil
	}
	// ]

	type_, err := ds.Type()
	if err != nil {
		if errors.Is(err, searchicon.HKVNotFound) {
			return false, nil
		} else {
			return false, err
		}
	}

	size, err := ds.Size()
	if err != nil {
		return false, err
	}

	switch type_ {
	case "Fixed":
		return size == iconsize, nil
	case "Scalable":
		minsize, err := ds.MinSize()
		if err != nil {
			return false, err
		}

		maxsize, err := ds.MaxSize()
		if err != nil {
			return false, err
		}

		return minsize <= iconsize && iconsize <= maxsize, nil
	case "Threshold":
		threshold := ds.Threshold()

		return size-threshold <= iconsize && iconsize <= size+threshold, nil
	}

	return false, nil
}

package searchicon

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var dirSizeMismatch = errors.New("theme subdirectory does not match the size constraint specified.")
var idxFormatErr = errors.New("index.theme file has an invalid format")

// Where idxFile is the index.theme file.
func dirMatchesSize(idxFile *os.File, subdir string, iconsize int, iconscale int) (bool, error) {
	// if Scale != iconscale [
	var scaleint int
	scalestr, err := getHKV(idxFile, subdir, "Scale")
	if err != nil {
		// default Scale to 1 as spec if not found
		scaleint = 1
	} else {
		scaleint, err = strconv.Atoi(scalestr)

		if err != nil {
			return false, fmt.Errorf(
				`failed to parse "Scale" of %v in file %v: %w: %w`,
				subdir, idxFile.Name(), err, idxFormatErr)
		}
	}

	if scaleint != iconscale {
		return false, nil
	}
	// ]

	typestr, err := getHKV(idxFile, subdir, "Type")
	if err != nil {
		return false, nil
	}

	sizestr, err := getHKV(idxFile, subdir, "Size")
	if err != nil {
		return false, fmt.Errorf(
			`missing "Size" key of %v in file %v: %w`,
			subdir, idxFile.Name(), idxFormatErr)
	}
	sizeint, err := strconv.Atoi(sizestr)
	if err != nil {
		return false, fmt.Errorf(
			`failed to parse "Size" of %v in file %v: %w: %w`,
			subdir, idxFile.Name(), err, idxFormatErr)
	}

	switch typestr {
	case "Fixed":
		return sizeint == iconsize, nil
	case "Scaled":
		var minint int
		minstr, err := getHKV(idxFile, subdir, "MinSize")
		if err != nil {
			minint = sizeint
		} else {
			minint, err = strconv.Atoi(minstr)
		}
		if err != nil {
			return false, fmt.Errorf(
				`failed to parse "MinSize" of %v in file %v: %w: %w`,
				subdir, idxFile.Name(), err, idxFormatErr)
		}

		var maxint int
		maxstr, err := getHKV(idxFile, subdir, "MaxSize")
		if err != nil {
			maxint = sizeint
		} else {
			maxint, err = strconv.Atoi(maxstr)
		}
		if err != nil {
			return false, fmt.Errorf(
				`failed to parse "MaxSize" of %v in file %v: %w: %w`,
				subdir, idxFile.Name(), err, idxFormatErr)
		}

		return minint <= iconsize && iconsize <= maxint, nil
	case "Threshold":
		var thint int
		thstr, err := getHKV(idxFile, subdir, "Threshold")
		if err != nil {
			thint = 2
		} else {
			thint, err = strconv.Atoi(thstr)
		}
		if err != nil {
			return false, fmt.Errorf(
				`failed to parse "Threshold" of %v in file %v: %w: %w`,
				subdir, idxFile.Name(), err, idxFormatErr)
		}

		return sizeint-thint <= iconsize && iconsize <= sizeint+thint, nil
	}

	return false, nil
}
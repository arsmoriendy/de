package searchicon

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

var dirSizeDistanceErr error = errors.New("Coultn't determine directory size distance")

func dirSizeDistance(idxFile *os.File, subdir string, iconsize int, iconscale int) (int, error) {
	// get Type [
	typestr, err := getHKV(idxFile.Name(), subdir, "Type")
	if err != nil {
		return 0, dirSizeDistanceErr
	}
	// ]

	// get Scale [
	var scaleint int

	scalestr, err := getHKV(idxFile.Name(), subdir, "Scale")
	if err == nil {
		scaleint, err = strconv.Atoi(scalestr)
	}

	if err != nil {
		scaleint = 1
	}
	// ]

	switch typestr {
	case "Fixed":
		// get Size [
		var sizeint int

		sizestr, err := getHKV(idxFile.Name(), subdir, "Size")
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}

		sizeint, err = strconv.Atoi(sizestr)
		if err != nil {
			err = fmt.Errorf("%w: %w", dirSizeDistanceErr, err)
			return 0, err
		}
		// ]

		dist := math.Abs(float64(scaleint*sizeint - iconsize*iconscale))
		return int(dist), nil
	case "Scalable", "Scaled":
		// get MinSize [
		var minsizeint int

		minsizestr, err := getHKV(idxFile.Name(), subdir, "MinSize")
		if err != nil {
			break
		}

		minsizeint, err = strconv.Atoi(minsizestr)
		if err != nil {
			err = fmt.Errorf("%w: %w: %w", dirSizeDistanceErr,
				errors.New("failed to convert MinSize to int"),
				err)
			return 0, err
		}
		// ]

		// get MaxSize [
		var maxsizeint int

		maxsizestr, err := getHKV(idxFile.Name(), subdir, "MaxSize")
		if err != nil {
			break
		}

		maxsizeint, err = strconv.Atoi(maxsizestr)
		if err != nil {
			err = fmt.Errorf("%w: %w: %w", dirSizeDistanceErr,
				errors.New("failed to convert MaxSize to int"),
				err)
			return 0, err
		}
		// ]

		if iconsize*iconscale < minsizeint*scaleint {
			return minsizeint*scaleint - iconsize*iconscale, nil
		}
		if iconsize*iconscale > maxsizeint*scaleint {
			return iconsize*iconscale - maxsizeint*scaleint, nil
		}
		return 0, nil
	}

	return 0, nil
}

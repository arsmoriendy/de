package searchicon

import (
	"os"
	"strconv"
)

func dirMatchesSize(idxFile *os.File, subdir string, iconsize int, iconscale int) bool {
	scalestr, err := getHKV(idxFile, subdir, "Scale")
	if err != nil {
		// TODO: handle err
	}
	scaleint, err := strconv.Atoi(scalestr)
	if err != nil {
		// TODO: handle err
	}

	sizestr, err := getHKV(idxFile, subdir, "Size")
	if err != nil {
		// TODO: handle err
	}
	sizeint, err := strconv.Atoi(sizestr)
	if err != nil {
		// TODO: handle err
	}

	minstr, err := getHKV(idxFile, subdir, "MinSize")
	if err != nil {
		// TODO: handle err
	}
	minint, err := strconv.Atoi(minstr)
	if err != nil {
		// TODO: handle err
	}

	maxstr, err := getHKV(idxFile, subdir, "MaxSize")
	if err != nil {
		// TODO: handle err
	}
	maxint, err := strconv.Atoi(maxstr)
	if err != nil {
		// TODO: handle err
	}

	thstr, err := getHKV(idxFile, subdir, "Threshhold")
	if err != nil {
		// TODO: handle err
	}
	thint, err := strconv.Atoi(thstr)
	if err != nil {
		// TODO: handle err
	}

	typestr, err := getHKV(idxFile, subdir, "Type")
	if err != nil {
		// TODO: handle err
	}

	// start matching
	if scaleint != iconscale {
		return false
	}

	if typestr == "Fixed" {
		return sizeint == iconsize
	}

	if typestr == "Scaled" {
		return minint <= iconsize && iconsize <= maxint
	}

	if typestr == "Threshhold" {
		return sizeint-thint <= iconsize && iconsize <= sizeint+thint
	}

	return false
}

package iconspec

import (
	"strconv"

	searchicon "github.com/arsmoriendy/de/pkgs/searchicon/gethkv"
)

type DirSpec struct {
	name string

	size      int
	scale     int
	context   string
	type_     string
	maxsize   int
	minsize   int
	threshold int

	// Is initialized vars for ints

	sizeIn      bool
	scaleIn     bool
	maxsizeIn   bool
	minsizeIn   bool
	thresholdIn bool

	idxFn string
}

func NewDir(idxFn string, name string) DirSpec {
	return DirSpec{
		name:        name,
		idxFn:       idxFn,
		sizeIn:      false,
		scaleIn:     false,
		maxsizeIn:   false,
		minsizeIn:   false,
		thresholdIn: false,
	}
}

func (d DirSpec) get(key string) (string, error) {
	return searchicon.GetHKV(d.idxFn, d.name, key)
}

func (d DirSpec) Size() (int, error) {
	if d.sizeIn {
		return d.size, nil
	}

	// get Size [
	var sizeint int

	sizestr, err := d.get("Size")
	if err != nil {
		return 0, err
	}

	sizeint, err = strconv.Atoi(sizestr)
	if err != nil {
		return 0, err
	}
	// ]

	d.sizeIn = true
	d.size = sizeint
	return sizeint, nil
}

func (d DirSpec) Scale() (int, error) {
	// get Size [
	var scaleint int

	scalestr, err := d.get("Scale")
	if err != nil {
		return 0, err
	}

	scaleint, err = strconv.Atoi(scalestr)
	if err != nil {
		return 0, err
	}
	// ]

	return idef(&(d.scale), (&d.scaleIn), scaleint), nil
}

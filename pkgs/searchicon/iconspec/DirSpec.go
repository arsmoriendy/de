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

// Wrapper for getting and initializing integer attributes.
// Parameter info:
// - a: attribute.
// - init: wether `a` is initialized.
// - key: key name in the index.theme file.
func (d DirSpec) initRetI(a *int, init *bool, key string) (int, error) {
	return initRet(a, *init, func() (int, error) {
		strval, err := d.get(key)
		if err != nil {
			return 0, err
		}

		intval, err := strconv.Atoi(strval)
		if err != nil {
			return 0, err
		}

		*init = true
		return intval, nil
	})
}

// Wrapper for getting and initializing string attributes.
// Uses the same parameters as [initRetI] (except for init).
func (d DirSpec) initRetS(a *string, key string) (string, error) {
	return initRet(a, *a == "", func() (string, error) { return d.get(key) })
}

func (d DirSpec) Size() (int, error) {
	return d.initRetI(&d.size, &d.sizeIn, "Size")
}

func (d DirSpec) Scale() (int, error) {
	scale, err := d.initRetI(&d.scale, &d.scaleIn, "Scale")
	if err != nil {
		scale = 1 // default
	}
	return scale, nil
}

func (d DirSpec) MaxSize() (int, error) {
	maxsize, err := d.initRetI(&d.maxsize, &d.maxsizeIn, "MaxSize")
	if err != nil {
		maxsize, err = d.Size()
	}
	return maxsize, err
}

func (d DirSpec) MinSize() (int, error) {
	minsize, err := d.initRetI(&d.maxsize, &d.maxsizeIn, "MinSize")
	if err != nil {
		minsize, err = d.Size()
	}
	return minsize, err
}

func (d DirSpec) Threshold() (int, error) {
	threshold, err := d.initRetI(&d.threshold, &d.thresholdIn, "Threshold")
	if err != nil {
		threshold = 1
	}
	return threshold, nil
}

func (d DirSpec) Context() (string, error) {
	return d.initRetS(&d.context, "Context")
}

func (d DirSpec) Type() (string, error) {
	return d.initRetS(&d.type_, "Type")
}

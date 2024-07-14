package iconspec

import "github.com/arsmoriendy/de/pkgs/searchicon/gethkv"

type IconSpec struct {
	name              string
	comment           string
	inherits          string
	directories       string
	scaledDirectories string
	example           string

	idxFn string
}

func New(idxFn string) IconSpec {
	return IconSpec{
		idxFn: idxFn,
	}
}

func (i IconSpec) get(key string) (string, error) {
	s, err := searchicon.GetHKV(i.idxFn, "Icon Theme", key)
	if err != nil {
		return "", err
	}

	return s, nil
}

// Wrapper for getting and initializing string attributes.
// Uses the same parameters as [DirSpec.initRetI] (except for init).
func (i IconSpec) initRetS(a *string, key string) (string, error) {
	return initRet(a, *a == "", func() (string, error) { return i.get(key) })
}

func (i IconSpec) Name() (string, error) {
	return i.initRetS(&i.name, "Name")
}

func (i IconSpec) Comment() (string, error) {
	return i.initRetS(&i.comment, "Comment")
}

func (i IconSpec) Inherits() (string, error) {
	return i.initRetS(&i.inherits, "Inherits")
}

func (i IconSpec) Directories() (string, error) {
	return i.initRetS(&i.directories, "Directories")
}

func (i IconSpec) ScaledDirectories() (string, error) {
	return i.initRetS(&i.scaledDirectories, "ScaledDirectories")
}

func (i IconSpec) Example() (string, error) {
	return i.initRetS(&i.example, "Example")
}

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

func (i IconSpec) Name() (string, error) {
	a := i.name
	init := a == ""
	key := "Name"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

func (i IconSpec) Comment() (string, error) {
	a := i.comment
	init := a == ""
	key := "Comment"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

func (i IconSpec) Inherits() (string, error) {
	a := i.inherits
	init := a == ""
	key := "Inherits"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

func (i IconSpec) Directories() (string, error) {
	a := i.directories
	init := a == ""
	key := "Directories"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

func (i IconSpec) ScaledDirectories() (string, error) {
	a := i.scaledDirectories
	init := a == ""
	key := "ScaledDirectories"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

func (i IconSpec) Example() (string, error) {
	a := i.example
	init := a == ""
	key := "Example"
	return initRet(&a, &init, func() (string, error) { return i.get(key) })
}

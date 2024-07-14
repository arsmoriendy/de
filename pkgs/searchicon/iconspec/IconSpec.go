package iconspec

import "github.com/arsmoriendy/de/pkgs/searchicon"

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
	name, err := i.get("Name")
	return sdef(&(i.name), name), err
}

func (i IconSpec) Comment() (string, error) {
	comment, err := i.get("Comment")
	return sdef(&(i.comment), comment), err
}

func (i IconSpec) Inherits() (string, error) {
	inherits, err := i.get("Inherits")
	return sdef(&(i.inherits), inherits), err
}

func (i IconSpec) Directories() (string, error) {
	directories, err := i.get("Directories")
	return sdef(&(i.directories), directories), err
}

func (i IconSpec) ScaledDirectories() (string, error) {
	scaledDirectories, err := i.get("ScaledDirectories")
	return sdef(&(i.scaledDirectories), scaledDirectories), err
}

func (i IconSpec) Example() (string, error) {
	example, err := i.get("Example")
	return sdef(&(i.example), example), err
}

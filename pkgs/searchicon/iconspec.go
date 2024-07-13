package searchicon

type iconspec struct {
	name              string
	comment           string
	inherits          string
	directories       string
	scaledDirectories string
	hidden            bool
	example           string

	idxFn string
}

func (i iconspec) getV(header string, key string) (string, error) {
	s, err := getHKV(i.idxFn, header, key)
	if err != nil {
		return "", err
	}

	return s, nil
}

func (i iconspec) Name() (string, error) {
	if i.name == "" {
		var err error
		i.name, err = i.getV("Icon Theme", "Name")
		return i.name, err
	}

	return i.name, nil
}

func (i iconspec) Comment() (string, error) {
	if i.comment != "" {
		return i.comment, nil
	}

	return i.getV("Icon Theme", "Comment")
}

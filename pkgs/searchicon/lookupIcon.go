package searchicon

func lookupIcon(iconname string, size string, scale string, theme string) string {
	rstring := ""

	idxFile, err := getThemeIdx(getThemeDirs(theme))
	if err != nil {
		// TODO: handle err
	}

	subdirs, err := getHKV(idxFile, "Icon Theme", "Directories")
	if err != nil {
		// TODO: handle err
	}

	walkDirsStr(subdirs, ',', func(dir string) error {
		return nil
	})

	return rstring
}

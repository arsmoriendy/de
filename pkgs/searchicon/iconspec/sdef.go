package iconspec

func sdef(s *string, def string) string {
	if *s != "" {
		return *s
	}

	*s = def
	return def
}


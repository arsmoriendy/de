package getde

func parseLine(line *string) (string, string) {
	key := ""
	value := ""
	lhs := true
	for _, c := range *line {
		if c == '=' {
			lhs = false
			continue
		}

		if lhs {
			key = key + string(c)
		} else {
			value = value + string(c)
		}
	}
	return key, value
}

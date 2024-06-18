package parsede

// replace "{key}" in format with the value of a key
func genFormatEntryFunc(format string) func(*map[string]string) string {
	return func(m *map[string]string) string {
		rstring := ""

		escaped := false
		capture := false
		key := ""
		for _, c := range format {
			if capture {
				switch c {
				case '}':
					rstring += (*m)[key]

					// reset
					capture = false
					key = ""
				default:
					key += string(c)
				}

				continue
			}

			if escaped {
				rstring += string(c)
				escaped = false
				continue
			}

			switch c {
			case '\\':
				escaped = true
			case '{':
				capture = true
			default:
				rstring += string(c)
			}

		}

		return rstring
	}
}

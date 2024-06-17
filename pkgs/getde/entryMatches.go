package getde

import "regexp"

func entryMatches(entry *map[string]string, regMap *map[string]*regexp.Regexp) bool {
	for k, reg := range *regMap {
		matches := reg.MatchString((*entry)[k])
		if !matches {
			return false
		}
	}

	return true
}

package getdefiles

import "regexp"

// Returns a map containing each entry's key to it's respective regex
func genRegMap(matchMap *map[string]string) map[string]*regexp.Regexp {
	rmap := map[string]*regexp.Regexp{}

	for k, v := range *matchMap {
		rmap[k] = regexp.MustCompile(v)
	}

	return rmap
}

package searchicon

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arsmoriendy/de/pkgs/parser"
)

// Gets a value from an index.theme file.
//
// HKV = Header Key Value.
func getHKV(idxFile *os.File, header string, key string) (string, error) {
	scanner := bufio.NewScanner(idxFile)

	// e = entry
	eheader := ""
	ekey := ""
	evalue := ""
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			continue
		}

		switch line[0] {
		case '[':
			eheader = getHeader(line)
			continue
		case '#':
			continue
		}

		if eheader != header {
			continue
		}

		ekey, evalue = parser.ParseLine(&line)

		if ekey != key {
			continue
		}

		return evalue, nil
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf(
		`Cannot find key "%v" in header "%v" in file "%v"`,
		key, header, idxFile.Name(),
	)
}

func getHeader(line string) string {
	rstring := ""

	for _, c := range line {
		switch c {
		case '[':
		case ']':
		default:
			rstring += string(c)
		}
	}

	return rstring
}

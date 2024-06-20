package searchicon

import (
	"bufio"
	"fmt"
	"os"

	"github.com/arsmoriendy/de/pkgs/parser"
)

var hkvNotFound = fmt.Errorf("could not find the value of specified header and key in the file")

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
		`with header %v, key %v, filename %v: %w`,
		header, key, idxFile.Name(), hkvNotFound,
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

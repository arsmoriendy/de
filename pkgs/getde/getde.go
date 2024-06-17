package getde

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetDe(entryfile os.File) string {
	rstring := ""

	scanner := bufio.NewScanner(&entryfile)

	entry := make(map[string]string)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 1 {
			continue
		}

		switch line[0] {
		case '[':
			if len(entry) != 0 {
				rstring = rstring + parseEntry(&entry) + "\n"
			}
			continue
		case '#':
			continue
		}

		key, value := parseLine(&line)

		switch key {
		case "Name":
		case "Icon":
		case "Exec":
		default:
			continue
		}

		entry[key] = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatalln(err)
	}

	rstring = rstring + parseEntry(&entry)

	return rstring
}

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

func parseEntry(entry *map[string]string) string {
	return fmt.Sprintf(
		"%v=%v=%v",
		(*entry)["Name"],
		(*entry)["Icon"],
		(*entry)["Exec"],
	)
}

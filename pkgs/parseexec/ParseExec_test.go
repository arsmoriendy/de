package parseexec

import (
	"testing"

	"github.com/arsmoriendy/de/pkgs/parser"
)

func TestLowerFMulti(t *testing.T) {
	entry := map[string]string{
		"Exec": "program %f -f",
	}
	opts := parser.Options{
		Names: []string{
			"/example/path/1",
			"/example/path/2",
			"/example/path/3",
		},
	}
	filename := ""

	exec := ParseExec(&entry, &opts, filename)
	exp := "program /example/path/1 -f"

	if exec != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

func FuzzLowerF(f *testing.F) {
	f.Add("/example/path")
	f.Add("/usr/local/bin/")
	f.Add("/root/ ")

	f.Fuzz(func(t *testing.T, file string) {
		entry := map[string]string{
			"Exec": "program %f -f",
		}
		opts := parser.Options{
			Names: []string{file},
		}
		filename := ""

		exec := ParseExec(&entry, &opts, filename)

		if exec != "program "+file+" -f" {
			t.Errorf("\nRaw: \t%v\nName: \t%v\nParsed: \t%v\n", entry["Exec"], file, exec)
		}
	})
}

type testParams struct {
	entry    map[string]string
	opts     parser.Options
	filename string
}

// Default testParams.
// Feel free to copy and override.
var defTP = testParams{
	entry: map[string]string{
		"Name":     "program-name",
		"Name[en]": "program-name-en",
		"Icon":     "program-icon",
		"Exec":     "program",
	},
	opts: parser.Options{
		Names: []string{
			"/dir/subdir/file",
			"/dir/subdir/file2",
		},
		Urls: []string{
			"proto://dest/path?param=foo",
			"proto://dest2/path?param=foo",
			"/dir/subdir/file", // file path can also be used
		},
	},
	filename: "program.desktop",
}

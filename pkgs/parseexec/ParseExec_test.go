package parseexec

import (
	"testing"

	"github.com/arsmoriendy/de/pkgs/parser"
)

// %f with zero filenames
func TestLowerFZero(t *testing.T) {
	param := defTP
	param.entry["Exec"] = "program %f -f"
	param.opts.Names = nil

	exec := param.call()
	exp := "program  -f"

	if exec != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

// %f with multiple filenames
func TestLowerFMulti(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %f -f"
	p.opts.Names = []string{
		"/example/path/1",
		"/example/path/2",
		"/example/path/3",
	}

	exec := p.call()
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
		p := defTP
		p.entry["Exec"] = "program %f -f"
		p.opts.Names = []string{file}

		exec := p.call()

		if exec != "program "+file+" -f" {
			t.Errorf("\nRaw: \t%v\nName: \t%v\nParsed: \t%v\n", p.entry["Exec"], file, exec)
		}
	})
}

func TestUpperF(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %F -f"

	exec := p.call()
	exp := "program /dir/subdir/file /dir/subdir/file2 -f"

	if exec != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

func TestUpperFZero(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %F -f"
	p.opts.Names = nil

	exec := p.call()
	exp := "program  -f"

	if exec != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

func TestI(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %i -f"

	exec := p.call()
	exp := `program --icon "program-icon" -f`

	if exec != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

func TestIZero(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %i -f"
	delete(p.entry, "Icon")

	exec := p.call()
	exp := `program  -f`

	if exec != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

// FC = field code, i.e. %_
func TestInvalidFC(t *testing.T) {
	p := defTP
	p.entry["Exec"] = "program %_ -f"

	exec := p.call()
	exp := "program  -f"

	if exec != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, exec)
	}
}

type testParams struct {
	entry    map[string]string
	opts     parser.Options
	filename string
}

func (tp testParams) call() string {
	return ParseExec(&tp.entry, &tp.opts, tp.filename)
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

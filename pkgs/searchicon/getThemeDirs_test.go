package searchicon

import (
	"path/filepath"
	"testing"
)

func TestGetThemeDirs(t *testing.T) {
	testBaseDir, err := filepath.Abs("../../test/data/sample_xdg_data_dir/icons/")
	if err != nil {
		t.Fatal(err)
	}

	tempBaseDirs := baseDirs
	baseDirs = testBaseDir

	// Test 1
	themeName := "SampleAdwaita"

	themeDirs := getThemeDirs(themeName)
	exp := filepath.Join(testBaseDir, themeName)

	if themeDirs != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, themeDirs)
	}

	// Test 2
	themeName = "Adwaita"

	themeDirs = getThemeDirs(themeName)
	exp = ""

	if themeDirs != exp {
		t.Fatalf("\nExpected:\t%v\nGot Result:\t%v", exp, themeDirs)
	}

	baseDirs = tempBaseDirs
}

package searchicon

import (
	"os"
	"testing"
)

func TestDirMatchesSize(t *testing.T) {
	idxFile, err := os.Open("../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme")
	if err != nil {
		t.Fatal(err)
	}

	// Test 1
	subdir := "scalable/devices"
	iconsize := 128
	iconscale := 1

	matches, err := dirMatchesSize(idxFile, subdir, iconsize, iconscale)
	if err != nil {
		t.Fatal(err)
	}

	if !matches {
		t.Fatal("Should match")
	}

	// Test 2: fixed type
	subdir = "16x16/ui"
	iconsize = 16

	matches, err = dirMatchesSize(idxFile, subdir, iconsize, iconscale)
	if err != nil {
		t.Fatal(err)
	}

	if !matches {
		t.Fatal("Should match")
	}

	// Test 3: fixed type, wrong size
	subdir = "16x16/ui"
	iconsize = 0

	matches, err = dirMatchesSize(idxFile, subdir, iconsize, iconscale)
	if err != nil {
		t.Fatal(err)
	}

	if matches {
		t.Fatal("Should not match")
	}

	// Test 4: scaled type
	subdir = "scalable/devices"
	iconsize = 8

	matches, err = dirMatchesSize(idxFile, subdir, iconsize, iconscale)
	if err != nil {
		t.Fatal(err)
	}

	if !matches {
		t.Fatal("Should match")
	}

	// Test 5: scaled type, wrong iconsize
	iconsize = 7

	matches, err = dirMatchesSize(idxFile, subdir, iconsize, iconscale)
	if err != nil {
		t.Fatal(err)
	}

	if matches {
		t.Fatal("Should not match")
	}

	// TODO: test threshold type
}

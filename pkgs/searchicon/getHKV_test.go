package searchicon

import (
	"os"
	"testing"
)

func TestGetHKV(t *testing.T) {
	idxFile, err := os.Open("../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme")
	header := "scalable/devices"
	key := "Type"

	if err != nil {
		t.Fatal(err)
	}

	value, err := getHKV(idxFile, header, key)
	exp := "Scalable"

	if err != nil {
		t.Fatal(err)
	}

	if value != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, value)
	}
}

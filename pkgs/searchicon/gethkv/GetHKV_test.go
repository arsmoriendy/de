package searchicon

import (
	"testing"
)

func TestGetHKV(t *testing.T) {
	// Test "Scale"
	fn := "../../test/data/sample_xdg_data_dir/icons/SampleAdwaita/index.theme"
	header := "scalable/devices"
	key := "Scale"

	value, err := GetHKV(fn, header, key)
	exp := ""

	if err == HKVNotFound {
		t.Errorf("should return error %v but doesn't", HKVNotFound)
	}

	if value != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, value)
	}

	// Test "Type"
	key = "Type"

	value, err = GetHKV(fn, header, key)
	exp = "Scalable"

	if err != nil {
		t.Fatal(err)
	}

	if value != exp {
		t.Errorf("\nExpected:\t%v\nGot Result:\t%v", exp, value)
	}
}

package searchicon

import "testing"

func TestAppendXdgDataDirs(t *testing.T) {
	xdgDataDirs := "/usr/local/share:/usr/share"
	s := "icons"

	result := appendXdgDataDirs(&xdgDataDirs, s)
	expected := "/usr/local/share/icons:/usr/share/icons"

	if result != expected {
		t.Fatalf(
			"\nExpected:\t%v\nGot Result:\t%v",
			expected,
			result,
		)
	}
}

package getdefiles

import "testing"

func TestSampleEntry(t *testing.T) {
	sampleEntry := map[string]string{
		"Name": "Firefox",
		"Icon": "firefox",
		"Exec": "firefox %u",
	}

	format := "{Name}={Icon}={Exec}"
	parseEntry := createParseEntry(format)
	result := parseEntry(&sampleEntry)

	expected := "Firefox=firefox=firefox %u"

	if result != expected {
		t.Fatalf(
			"\nExpected:\t%v\nGot Result:\t%v",
			expected,
			result,
		)
	}
}

func TestEscapedFormat(t *testing.T) {
	sampleEntry := map[string]string{
		"Name": "Firefox",
		"Icon": "firefox",
		"Exec": "firefox %u",
	}

	escapedFormat := "{Name}={Icon}=\\{Exec}"
	parseEntry := createParseEntry(escapedFormat)
	result := parseEntry(&sampleEntry)

	expected := "Firefox=firefox={Exec}"

	if result != expected {
		t.Fatalf(
			"\nExpected:\t%v\nGot Result:\t%v",
			expected,
			result,
		)
	}
}

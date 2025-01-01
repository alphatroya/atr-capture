package bookmarks

import (
	"testing"
)

func TestTitleDetection(t *testing.T) {
	tests := []struct {
		in string
		ex bool
	}{
		{
			in: "https://tanaschita.com/ios-universal-links-swiftui/",
			ex: true,
		},
	}

	for _, test := range tests {
		r := containsHTTPLink(test.in)
		if r != test.ex {
			t.Errorf("containsHTTPLink(%q) = %v; want %v", test.in, r, test.ex)
		}
	}
}
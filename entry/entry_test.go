package entry

import (
	"testing"
	"time"
)

func TestEntryResultBuilding(t *testing.T) {
	year, month, day := 2023, time.October, 1
	hour, minute, second := 9, 45, 0
	time := time.Date(year, month, day, hour, minute, second, 0, time.UTC)

	tests := []struct {
		in       string
		taglist  []string
		expected string
	}{
		{
			in:       "123",
			expected: "- **09:45** 123",
		},
		{
			in: `112
123`,
			expected: `- **09:45** 112
  123`,
		},
		{
			in:       "abc",
			taglist:  []string{"a", "b"},
			expected: "- **09:45** abc #a #b",
		},
		{
			in:       "abc",
			taglist:  []string{"a", "todo"},
			expected: "- TODO **09:45** abc #a",
		},
	}

	for _, test := range tests {
		result := NewEntry(test.in, test.taglist).Build(time)
		if result != test.expected {
			t.Errorf("Build(%s) = \"%s\"; want \"%s\"", test.in, result, test.expected)
		}
	}
}

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
		expected string
	}{
		{
			in:       "123",
			expected: "- **9:45** 123",
		},
		{
			in: `112
123`,
			expected: `- **9:45** 112
  123`,
		},
	}

	for _, test := range tests {
		result := NewEntry(test.in, []string{}).Build(time)
		if result != test.expected {
			t.Errorf("Build(%s) = \"%s\"; want \"%s\"", test.in, result, test.expected)
		}
	}
}

package entry

import "testing"

func TestEntryResultBuilding(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{
			in:       "123",
			expected: "- 123",
		},
		{
			in: `112
123`,
			expected: `- 112
  123`,
		},
	}

	for _, test := range tests {
		result := NewEntry(test.in, []string{}).Build(false)
		if result != test.expected {
			t.Errorf("Build(%s) = \"%s\"; want \"%s\"", test.in, result, test.expected)
		}
	}
}

package quote

import (
	"testing"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func TestFormatQuote(t *testing.T) {
	tests := []struct {
		name     string
		in       string
		taglist  []string
		expected string
	}{
		{
			name:     "not a quote tag",
			in:       "test quote",
			taglist:  []string{},
			expected: "test quote",
		},
		{
			name:    "simple quote",
			in:      "test quote",
			taglist: []string{"quote"},
			expected: `
  > test quote
`,
		},
		{
			name:    "multiline quote",
			in:      "test quote\ntest2",
			taglist: []string{"quote"},
			expected: `
  > test quote
  > test2
`,
		},
		{
			name:    "quote with author",
			in:      "test quote\ntest2\n\nauthor",
			taglist: []string{"quote"},
			expected: `
  > test quote
  > test2
  >
  > _author_
`,
		},
	}

	for _, test := range tests {
		result := FormatQuoteIfNeeded(draft.Draft{
			Text: test.in,
			Tags: test.taglist,
		})
		if result.Text != test.expected {
			t.Errorf("%s: Build(%s) = \"%s\"; want \"%s\"", test.name, test.in, result.Text, test.expected)
		}
	}
}

package save

import (
	"testing"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func TestBuildNote(t *testing.T) {
	tests := []struct {
		name     string
		draft    draft.Draft
		expected string
	}{
		{
			name: "Without TODO",
			draft: draft.Draft{
				Text:   "This is a note.",
				IsTODO: false,
			},
			expected: "- This is a note.",
		},
		{
			name: "With TODO",
			draft: draft.Draft{
				Text:   "This is a TODO item.",
				IsTODO: true,
			},
			expected: "- TODO This is a TODO item.",
		},
		{
			name: "With Multi-Line Text",
			draft: draft.Draft{
				Text:   "First line.\nSecond line.",
				IsTODO: false,
			},
			expected: "- First line.\n  Second line.",
		},
		{
			name: "With Multi-Line Text and TODO",
			draft: draft.Draft{
				Text:   "First line TODO.\nSecond line TODO.",
				IsTODO: true,
			},
			expected: "- TODO First line TODO.\n  Second line TODO.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result, want := buildNote(tt.draft), tt.expected; want != result {
				t.Errorf("TestBuildNote(%q) = \"%v\"; want \"%v\"", tt.draft.Text, result, want)
			}
		})
	}
}

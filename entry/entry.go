package entry

import (
	"fmt"
	"strings"
)

const dashPrefix = "- "

type Entry struct {
	text string
}

func NewEntry(text string) Entry {
	return Entry{
		text: text,
	}
}

func (e Entry) Build() string {
	return fmt.Sprintf("%s%s", dashPrefix, padTextExceptFirstLine(e.text))
}

func padTextExceptFirstLine(text string) string {
	lines := strings.Split(text, "\n")
	padding := len(dashPrefix)
	for i, line := range lines {
		if i == 0 {
			continue
		}
		lines[i] = strings.Repeat(" ", padding) + line
	}
	return strings.Join(lines, "\n")
}

package entry

import (
	"fmt"
	"strings"
	"time"
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
	currentTime := time.Now()
	formattedTime := fmt.Sprintf("%d:%02d", currentTime.Hour(), currentTime.Minute())
	return fmt.Sprintf("%s**%s** %s", dashPrefix, formattedTime, padTextExceptFirstLine(e.text))
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

package entry

import (
	"fmt"
	"strings"
	"time"
)

const (
	dashPrefix = "- "
	todoMark   = "TODO "
)

type Entry struct {
	text string
	tags []string
}

func NewEntry(text string, tags []string) Entry {
	return Entry{
		text: text,
		tags: tags,
	}
}

func (e Entry) Build(time time.Time) string {
	formattedTime := fmt.Sprintf("%02d:%02d", time.Hour(), time.Minute())

	t := ""
	tagslist := ""
	for _, tag := range e.tags {
		if tag == "todo" {
			t = todoMark
			continue
		}
		tagslist += "#" + tag + " "
	}

	tagslist = strings.TrimSpace(tagslist)
	return fmt.Sprintf("%s%s**%s** %s", dashPrefix, t, formattedTime, padTextExceptFirstLine(e.text, tagslist))
}

func padTextExceptFirstLine(text string, tagslist string) string {
	lines := strings.Split(text, "\n")
	padding := len(dashPrefix)
	for i, line := range lines {
		if i == 0 {
			if tagslist != "" {
				lines[i] = fmt.Sprintf("%s %s", line, tagslist)
			}
			continue
		}
		lines[i] = strings.Repeat(" ", padding) + line
	}
	return strings.Join(lines, "\n")
}

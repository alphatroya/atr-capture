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
	text    string
	tags    []string
	content string
}

func NewEntry(text string, tags []string, content string) Entry {
	return Entry{
		text:    text,
		tags:    tags,
		content: content,
	}
}

func (e Entry) Build(time time.Time) string {
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
	result := fmt.Sprintf("%s%s %s", dashPrefix, t, padTextExceptFirstLine(e.text, tagslist))
	if e.content == "" {
		return result
	}
	return fmt.Sprintf("%s\n\n---\n%s", result, e.content)
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

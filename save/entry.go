package save

import (
	"fmt"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

const (
	dashPrefix = "- "
	todoMark   = "TODO "
)

func buildNote(d draft.Draft) string {
	t := ""
	tagslist := ""
	for _, tag := range d.Tags {
		if tag == "todo" {
			t = todoMark
			continue
		}
		tagslist += "#" + tag + " "
	}

	tagslist = strings.TrimSpace(tagslist)
	return fmt.Sprintf("%s%s %s\n", dashPrefix, t, padTextExceptFirstLine(d.Text, tagslist))
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

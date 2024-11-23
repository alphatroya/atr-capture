package save

import (
	"fmt"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

const (
	dashPrefix         = "- "
	todoMark           = "TODO "
	bookmarkNoteSuffix = "_bookmark"
)

func buildNote(d draft.Draft, noteTitle string) string {
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
	result := fmt.Sprintf("%s%s %s", dashPrefix, t, padTextExceptFirstLine(d.Text, tagslist))
	if d.Post == nil || !d.Post.IsContentAvailable() {
		return result
	}
	return fmt.Sprintf("%s\n\n[[%s%s]]", result, noteTitle, bookmarkNoteSuffix)
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

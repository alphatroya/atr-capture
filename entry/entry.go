package entry

import (
	"fmt"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

const (
	dashPrefix = "- "
	todoMark   = "TODO "
)

func Build(d draft.Draft) string {
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
	if d.Post == nil || d.Post.Title == "" {
		return result
	}
	return fmt.Sprintf("%s\n\n[[%s]]", result, d.Post.Title)
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

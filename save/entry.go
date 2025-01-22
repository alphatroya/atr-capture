package save

import (
	"fmt"
	"strings"

	"github.com/alphatroya/atr-capture/draft"
)

const (
	dashPrefix = "- "
	todoMark   = "TODO "
)

func buildNote(d draft.Draft) string {
	t := ""
	if d.IsTODO {
		t = todoMark
	}
	return fmt.Sprintf("%s%s%s", dashPrefix, t, padTextExceptFirstLine(d.Text))
}

func padTextExceptFirstLine(text string) string {
	lines := strings.Split(text, "\n")
	padding := len(dashPrefix)
	for i, line := range lines {
		if i == 0 {
			lines[i] = line
			continue
		}
		lines[i] = strings.Repeat(" ", padding) + line
	}
	return strings.Join(lines, "\n")
}

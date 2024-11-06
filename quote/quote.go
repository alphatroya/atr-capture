package quote

import (
	"fmt"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func FormatQuoteIfNeeded(d draft.Draft) draft.Draft {
	for _, tag := range d.Tags {
		if tag == "quote" {
			d = draft.Draft{
				Text: fmt.Sprintf("\n%s\n", padQuote(d.Text)),
				Tags: d.Tags,
			}
			return d
		}
	}
	return d
}

func padQuote(text string) string {
	lines := strings.Split(text, "\n")
	var authorMode bool
	if len(lines) > 3 && lines[len(lines)-2] == "" {
		authorMode = true
	}

	for i, line := range lines {
		if i == len(lines)-1 && authorMode {
			lines[i] = "  > " + "_" + line + "_"
		} else {
			if line == "" {
				lines[i] = "  >"
			} else {
				lines[i] = "  > " + line
			}
		}
	}
	return strings.Join(lines, "\n")
}

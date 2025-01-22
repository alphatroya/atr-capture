package bookmarks

import (
	"fmt"
	"regexp"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func containsHTTPLink(s string) bool {
	re := regexp.MustCompile(`^http[s]?://[^\s]+`)
	return re.MatchString(s)
}

func ExtractAndFormatLinkTitles(text string) (draft.Draft, error) {
	lines := strings.Split(text, "\n")
	linesResult := make([]string, 0, len(lines))
	d := draft.Draft{}
	for _, line := range lines {
		fragments := strings.Split(line, " ")

		fragmentsResult := make([]string, 0, len(fragments))
		for _, fragment := range fragments {
			if fragment != "" && containsHTTPLink(fragment) {
				p, err := requestPageContent(fragment)
				if err == nil {
					d.Post = p
					fragment = fmt.Sprintf("[%s](%s)", d.Post.Title, d.Post.URL)
				}
			}

			fragmentsResult = append(fragmentsResult, fragment)
		}

		linesResult = append(linesResult, strings.Join(fragmentsResult, " "))
	}
	d.Text = strings.Join(linesResult, "\n")
	return d, nil
}

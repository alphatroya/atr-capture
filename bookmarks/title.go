package bookmarks

import (
	"fmt"
	"regexp"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func containsHTTPLink(s string) bool {
	re := regexp.MustCompile(`http[s]?://[^\s]+`)
	return re.MatchString(s)
}

func RequestTitleIfNeeded(d draft.Draft) (draft.Draft, bool, error) {
	lines := strings.Split(d.Text, "\n")
	linesResult := make([]string, 0, len(lines))
	for _, line := range lines {
		fragments := strings.Split(line, " ")

		fragmentsResult := make([]string, 0, len(fragments))
		for _, fragment := range fragments {
			if len(fragment) != 0 && containsHTTPLink(fragment) {
				d.Post = &draft.Post{
					URL: fragment,
				}
				d, err := requestPage(d)
				if err == nil {
					fragment = fmt.Sprintf("[%s](%s)", d.Post.Title, d.Post.URL)
				}
			}

			fragmentsResult = append(fragmentsResult, fragment)
		}

		linesResult = append(linesResult, strings.Join(fragmentsResult, " "))
	}
	return draft.Draft{
		Text: strings.Join(linesResult, "\n"),
		Tags: d.Tags,
		Post: d.Post,
	}, d.Post != nil, nil
}

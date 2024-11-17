package bookmarks

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	"golang.org/x/net/html"
)

func containsHTTPLink(s string) bool {
	re := regexp.MustCompile(`http[s]?://[^\s]+`)
	return re.MatchString(s)
}

func RequestTitleIfNeeded(d draft.Draft) (draft.Draft, error) {
	fragments := strings.Split(d.Text, " ")

	var url string
	var results []string
	for _, fragment := range fragments {
		if len(fragment) != 0 && containsHTTPLink(fragment) {
			t, err := fetchTitle(fragment)
			if err == nil {
				url = fragment
				fragment = fmt.Sprintf("[%s](%s)", t, fragment)
			}
		}

		results = append(results, fragment)
	}
	return draft.Draft{
		Text: strings.Join(results, " "),
		Tags: d.Tags,
		URL:  url,
	}, nil
}

// FetchTitle fetches the HTML title of the page at the given URL
func fetchTitle(url string) (string, error) {
	// Send an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse the HTML response
	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			// End of the document, return an error if we didn't find a title
			return "", fmt.Errorf("no title element found")
		case html.StartTagToken:
			// Get the tag name
			t := z.Token()
			if t.Data == "title" {
				// Read the text within the title tag
				z.Next()
				return strings.TrimSpace(z.Token().Data), nil
			}
		}
	}
}

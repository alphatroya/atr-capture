package title

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// FetchTitle fetches the HTML title of the page at the given URL
func FetchTitle(url string) (string, error) {
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

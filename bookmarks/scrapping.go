package bookmarks

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	htmltomarkdown "github.com/JohannesKaufmann/html-to-markdown/v2"
)

func RequestPage(d draft.Draft) (draft.Draft, error) {
	if d.URL == "" {
		return d, nil
	}

	resp, err := http.Get(d.URL)
	if err != nil {
		return d, fmt.Errorf("Can't fetch page, url=%s, err=%w", d.URL, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return d, fmt.Errorf("Error: received non-200 response status: %s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return d, err
	}

	input := string(body)
	markdown, err := htmltomarkdown.ConvertString(input)
	if err != nil {
		return d, fmt.Errorf("Can't parse url, url=%s, err=%w", d.URL, err)
	}

	d.Content = markdown
	return d, nil
}

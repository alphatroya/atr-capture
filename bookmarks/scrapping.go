package bookmarks

import (
	"fmt"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	readability "github.com/go-shiori/go-readability"
)

func RequestPage(d draft.Draft) (draft.Draft, error) {
	if d.URL == "" {
		return d, nil
	}

	article, err := readability.FromURL(d.URL, 30*time.Second)
	if err != nil {
		return d, fmt.Errorf("failed to get data content, url=%s, %w", d.URL, err)
	}

	d.Content = article.Content
	return d, nil
}

package bookmarks

import (
	"fmt"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	readability "github.com/go-shiori/go-readability"
)

func requestPage(d draft.Draft) (draft.Draft, error) {
	if d.Post == nil {
		return d, nil
	}

	article, err := readability.FromURL(d.Post.URL, 30*time.Second)
	if err != nil {
		return d, fmt.Errorf("failed to get data content, url=%s, %w", d.Post.URL, err)
	}

	d.Post.Title = article.Title
	d.Post.Content = article.Content
	return d, nil
}

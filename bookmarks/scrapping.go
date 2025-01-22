package bookmarks

import (
	"fmt"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	readability "github.com/go-shiori/go-readability"
)

func requestPageContent(url string) (*draft.Post, error) {
	article, err := readability.FromURL(url, 30*time.Second)
	if err != nil {
		return nil, fmt.Errorf("requestPageContent: failed to get data content, url=%s, err=%w", url, err)
	}

	if article.Title == "" {
		return nil, fmt.Errorf("requestPageContent: failed to get page title, url=%s, err=%w", url, err)
	}
	p := new(draft.Post)
	p.URL = url
	p.Title = article.Title
	p.Content = article.Content
	return p, nil
}

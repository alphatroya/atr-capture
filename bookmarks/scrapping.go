package bookmarks

import (
	"context"
	"fmt"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	"github.com/charmbracelet/huh/spinner"
	readability "github.com/go-shiori/go-readability"
)

func requestPageContent(url string) (*draft.Post, error) {
	pch := make(chan *draft.Post)
	errCh := make(chan error)
	ctx, done := context.WithCancel(context.Background())
	go func() {
		defer done()
		article, err := readability.FromURL(url, 30*time.Second)
		if err != nil {
			errCh <- fmt.Errorf("requestPageContent: failed to get data content, url=%s, err=%w", url, err)
			return
		}

		if article.Title == "" {
			errCh <- fmt.Errorf("requestPageContent: failed to get page title, url=%s, err=%w", url, err)
			return
		}
		p := new(draft.Post)
		p.URL = url
		p.Title = article.Title
		p.Content = article.Content
		pch <- p
		return
	}()

	go func() {
		spinner.New().
			Type(spinner.Dots).
			Title("Requesting page content...").
			Context(ctx).
			Run()
	}()

	select {
	case p := <-pch:
		return p, nil
	case err := <-errCh:
		return nil, err
	}
}

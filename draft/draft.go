package draft

import (
	"fmt"
	"os"
	"path/filepath"
)

var draftLocation string

func init() {
	config, err := os.UserConfigDir()
	if err != nil {
		panic(fmt.Sprintf("can't locate user config location: %s", err))
	}

	draftLocation = filepath.Join(config, "atr-capture", "draft.json")
	draftDir := filepath.Dir(draftLocation)

	if _, err := os.Stat(draftDir); os.IsNotExist(err) {
		err := os.MkdirAll(draftDir, os.ModePerm)
		if err != nil {
			panic(fmt.Sprintf("can't create directory %s: %s", draftDir, err))
		}
	}
}

type Draft struct {
	Text   string `json:"text"`
	Post   *Post  `json:"post"`
	IsTODO bool   `json:"isTodo"`
}

func (d Draft) ContainURL() bool {
	return d.Post != nil
}

type Post struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `jsong:"content"`
}

func (p *Post) IsContentAvailable() bool {
	return p.Title != "" && p.Content != ""
}

func (d Draft) IsEmpty() bool {
	return len(d.Text) == 0
}

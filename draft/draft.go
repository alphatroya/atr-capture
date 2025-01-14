package draft

import (
	"encoding/json"
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

func (d Draft) SaveIfNeeded() (bool, error) {
	if d.IsEmpty() {
		return false, nil
	}

	data, err := json.Marshal(d)
	if err != nil {
		return false, fmt.Errorf("error marshaling draft to JSON: %w", err)
	}

	file, err := os.Create(draftLocation)
	if err != nil {
		return false, fmt.Errorf("error creating/opening file at %s: %w", draftLocation, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return false, fmt.Errorf("error writing to file: %w", err)
	}

	return true, nil
}

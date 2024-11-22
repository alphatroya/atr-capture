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
	Text string   `json:"text"`
	Tags []string `json:"tags"`
	Post *Post    `json:"post"`
}

type Post struct {
	URL     string `json:"url"`
	Title   string `json:"title"`
	Content string `jsong:"content"`
}

func (d Draft) IsEmpty() bool {
	return len(d.Text) == 0 && len(d.Tags) == 0
}

func SaveDraft(draft Draft) error {
	data, err := json.Marshal(draft)
	if err != nil {
		return fmt.Errorf("error marshaling draft to JSON: %w", err)
	}

	file, err := os.Create(draftLocation)
	if err != nil {
		return fmt.Errorf("error creating/opening file at %s: %w", draftLocation, err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func DropDraft() {
	if _, err := os.Stat(draftLocation); err == nil {
		err = os.Remove(draftLocation)
		if err != nil {
			fmt.Printf("error removing draft file at %s: %v\n", draftLocation, err)
		} else {
			fmt.Println("Draft file successfully removed.")
		}
	}
}

func RestoreDraft() (Draft, bool) {
	if _, err := os.Stat(draftLocation); os.IsNotExist(err) {
		return Draft{}, false
	}

	data, err := os.ReadFile(draftLocation)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		return Draft{}, false
	}

	var draft Draft
	if err := json.Unmarshal(data, &draft); err != nil {
		fmt.Printf("error unmarshaling JSON to Draft: %v\n", err)
		return Draft{}, false
	}

	return draft, true
}

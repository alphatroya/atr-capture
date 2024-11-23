package save

import (
	"fmt"
	"os"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/generator"
)

var envs env.Envs

func init() {
	var err error
	envs, err = env.CheckEnvs()
	if err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}
}

func SaveToJournal(nt string) error {
	file, err := os.OpenFile(envs.TodayJournalPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("Error opening the journal file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + fmt.Sprintf("- {{embed [[%s]]}}", nt))
	return err
}

func SaveToPages(d draft.Draft) (string, error) {
	noteTitle := generator.GenerateQuickNoteTitle(time.Now())
	file, err := os.OpenFile(envs.PagesPath()+noteTitle+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return noteTitle, fmt.Errorf("Error opening the page file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(buildNote(d, noteTitle))
	if err != nil {
		return noteTitle, fmt.Errorf("Error writing the page file: %w", err)
	}

	if d.Post != nil && d.Post.IsContentAvailable() {
		err = SaveToPagesContent(d)
	}

	return noteTitle, err
}

func SaveToPagesContent(d draft.Draft) error {
	file, err := os.OpenFile(envs.PagesPath()+d.Post.Title+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("Error opening the page content file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(d.Post.Content)
	return err
}

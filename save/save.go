package save

import (
	"fmt"
	"os"
	"path/filepath"
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
		fmt.Printf("error in configuration: unable to check environment variables: %s\n", err)
		os.Exit(1)
	}
}

func SaveToJournal(nt string) error {
	file, err := os.OpenFile(envs.TodayJournalPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening the journal file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + fmt.Sprintf("- {{embed [[%s]]}}", nt))
	return err
}

func SaveToPages(d draft.Draft, saveContent bool) (string, error) {
	noteTitle := generator.GenerateQuickNoteTitle(time.Now())
	file, err := os.OpenFile(envs.PagesPath()+noteTitle+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return noteTitle, fmt.Errorf("error opening the page file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(buildNote(d))
	if err != nil {
		return noteTitle, fmt.Errorf("error writing the page file: %w", err)
	}

	if saveContent {
		err = SaveToPagesContent(d)
	}

	return noteTitle, err
}

func SaveToPagesContent(d draft.Draft) error {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory, %w", err)
	}
	path := filepath.Join(dirname, "Downloads", d.Post.Title+".html")
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening the page content file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(d.Post.Content)
	return err
}

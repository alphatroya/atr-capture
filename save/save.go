package save

import (
	"fmt"
	"os"
	"path/filepath"

	"git.sr.ht/~alphatroya/atr-capture/draft"
)

func SaveToJournal(nt string, path string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening the journal file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + fmt.Sprintf("- {{embed [[%s]]}}", nt))
	return err
}

func SaveToPages(notePath string, d draft.Draft, saveContent bool) error {
	file, err := os.OpenFile(notePath, os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("error opening the page file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString(buildNote(d))
	if err != nil {
		return fmt.Errorf("error writing the page file: %w", err)
	}

	if saveContent {
		err = saveToDownloadsPageContent(d)
	}

	return err
}

func saveToDownloadsPageContent(d draft.Draft) error {
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

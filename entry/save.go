package entry

import (
	"fmt"
	"os"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/generator"
)

func SaveToJournal(nt string) error {
	envs, err := env.CheckEnvs()
	if err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}

	file, err := os.OpenFile(envs.TodayJournalPath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString("\n" + fmt.Sprintf("- {{embed [[%s]]}}", nt))
	return err
}

func SaveToPages(out string) (string, error) {
	envs, err := env.CheckEnvs()
	if err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}

	noteTitle := generator.GenerateQuickNoteTitle(time.Now())

	file, err := os.OpenFile(envs.PagesPath()+noteTitle+".md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.WriteString(out)
	return noteTitle, err
}

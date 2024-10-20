package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/entry"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"github.com/charmbracelet/huh"
)

func main() {
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

	var text string
	var tags []string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Capture this").
				ShowLineNumbers(true).
				Validate(func(in string) error {
					in = strings.TrimSpace(in)
					if len(in) == 0 {
						return errors.New("quick capture text can't be empty")
					}
					return nil
				}).
				Value(&text),

			huh.NewMultiSelect[string]().
				Title("Select tags").
				Options(
					huh.NewOption("TODO", "todo"),
					huh.NewOption("📚 Book to read", "books"),
					huh.NewOption("🛍️ Book to buy", "books-to-buy"),
					huh.NewOption("🍿 Movie", "movies"),
					huh.NewOption("🤔 Ideas", "ideas"),
				).
				Value(&tags),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Error filling the form:", err)
		os.Exit(1)
	}
	out := entry.NewEntry(text, tags).Build()

	_, err = file.WriteString("\n" + out)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		os.Exit(1)
	}

	fmt.Println("Text appended successfully!")
}

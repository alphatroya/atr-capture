package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"git.sr.ht/~alphatroya/atr-capture/bookmarks"
	"git.sr.ht/~alphatroya/atr-capture/draft"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/forms"
	"git.sr.ht/~alphatroya/atr-capture/quote"
	"git.sr.ht/~alphatroya/atr-capture/save"
	"github.com/charmbracelet/huh"
)

func main() {
	if _, err := env.CheckEnvs(); err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}

	d := draft.RestoreOrNewDraft(
		func() bool {
			confirm, err := forms.ConfirmRestoreDraftDialog()
			if err != nil {
				fmt.Println("Error filling the form: ", err)
				os.Exit(1)
			}
			return confirm
		},
	)

	var isTodo bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Enter your note ✍️").
				ShowLineNumbers(true).
				Validate(func(in string) error {
					in = strings.TrimSpace(in)
					if len(in) == 0 {
						return errors.New("quick capture text can't be empty")
					}
					return nil
				}).
				Value(&d.Text),

			huh.NewMultiSelect[string]().
				Title("Select tags").
				Options(
					huh.NewOption("📚 Book to read", "books"),
					huh.NewOption("🛍️ Book to buy", "books-to-buy"),
					huh.NewOption("🍿 Movie", "movies"),
					huh.NewOption("🤔 Ideas", "ideas"),
					huh.NewOption("✍️ Blog", "blog"),
					huh.NewOption("💬 Quote", "quote"),
					huh.NewOption("🏃‍♂️ Running", "running"),
				).
				Value(&d.Tags),

			huh.NewConfirm().
				Title("Mark your note as TODO?").
				Value(&isTodo),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Error filling the form:", err)
		saveDraftIfNeeded(d)
		os.Exit(1)
	}

	if isTodo {
		d.Tags = append(d.Tags, "todo")
	}

	d = quote.FormatQuoteIfNeeded(d)
	d, containURL, err := bookmarks.RequestTitleIfNeeded(d)
	if err != nil {
		fmt.Println("Error requesting page title: ", err)
		saveDraftIfNeeded(d)
		os.Exit(1)
	}

	saveContent := false
	if containURL {
		saveContent = forms.RequestSavingContent()
		if err != nil {
			fmt.Println("Error requesting page content: ", err)
			saveDraftIfNeeded(d)
			os.Exit(1)
		}
	}

	nt, err := save.SaveToPages(d, saveContent)
	if err == nil {
		err = save.SaveToJournal(nt)
	}

	if err != nil {
		fmt.Println("Error writing to the file: ", err)
		saveDraftIfNeeded(d)
		os.Exit(1)
	}

	draft.DropDraft()
	fmt.Printf("Quick capture saved, a new note created: %s.md\n", nt)
}

func saveDraftIfNeeded(d draft.Draft) {
	saved, err := d.SaveIfNeeded()
	if err != nil {
		fmt.Println("Error saving the draft: ", err)
	} else if saved {
		fmt.Println("Draft saved for future use")
	}
}

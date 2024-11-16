package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/draft"
	"git.sr.ht/~alphatroya/atr-capture/entry"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/quote"
	"git.sr.ht/~alphatroya/atr-capture/title"
	"github.com/charmbracelet/huh"
)

func main() {
	_, err := env.CheckEnvs()
	if err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}

	d := newDraft()
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

	if err = form.Run(); err != nil {
		fmt.Println("Error filling the form:", err)
		saveDraftIfNeeded(d)

		os.Exit(1)
	}

	if isTodo {
		d.Tags = append(d.Tags, "todo")
	}

	d = quote.FormatQuoteIfNeeded(d)
	d, err = title.RequestTitleIfNeeded(d)
	if err != nil {
		fmt.Println("Error requesting page title: ", err)
		saveDraftIfNeeded(d)
		os.Exit(1)
	}

	out := entry.NewEntry(d.Text, d.Tags).Build(time.Now())
	nt, err := entry.SaveToPages(out)
	if err == nil {
		err = entry.SaveToJournal(nt)
	}

	if err != nil {
		fmt.Println("Error writing to the file: ", err)
		saveDraftIfNeeded(d)
		os.Exit(1)
	}

	draft.DropDraft()
	fmt.Println("The text has been successfully appended!")
}

func saveDraftIfNeeded(d draft.Draft) {
	if !d.IsEmpty() {
		if err := draft.SaveDraft(d); err != nil {
			fmt.Println("Error saving the draft: ", err)
		} else {
			fmt.Println("Draft saved for future use")
		}
	}
}

func newDraft() draft.Draft {
	d, restored := draft.RestoreDraft()
	if !restored {
		return d
	}

	usePrevDraft := true
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Found a draft, use it?").
				Value(&usePrevDraft),
		),
	)

	if err := form.Run(); err != nil {
		fmt.Println("Error filling the form:", err)
		os.Exit(1)
	}

	if usePrevDraft {
		return d
	}
	draft.DropDraft()
	return draft.Draft{}
}

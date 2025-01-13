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
	"git.sr.ht/~alphatroya/atr-capture/save"
	"git.sr.ht/~alphatroya/atr-capture/tags"
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

	err := huh.NewText().
		Title("Enter your note ✍️").
		ShowLineNumbers(true).
		Validate(func(in string) error {
			in = strings.TrimSpace(in)
			if len(in) == 0 {
				return errors.New("quick capture text can't be empty")
			}
			return nil
		}).
		Value(&d.Text).
		Run()
	checkErr("Form aborted: ", err, d)

	d, containURL, err := bookmarks.RequestTitleIfNeeded(d)
	d.Tags, err = forms.PickUpTags(tags.RequestTags())
	checkErr("Form aborted: ", err, d)

	var isTodo bool
	if huh.NewConfirm().
		Title("Mark your note as TODO?").
		Value(&isTodo).
		Run(); isTodo {
		d.Tags = append(d.Tags, "todo")
	}
	checkErr("Form aborted: ", err, d)

	saveContent := false
	if containURL {
		saveContent = forms.RequestSavingContent()
		checkErr("Error requesting page content: ", err, d)
	}

	nt, err := save.SaveToPages(d, saveContent)
	if err == nil {
		err = save.SaveToJournal(nt)
	}

	checkErr("Error writing to the file: ", err, d)

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

func checkErr(message string, err error, saveDraft draft.Draft) {
	if err != nil {
		fmt.Println(message, err)
		saveDraftIfNeeded(saveDraft)
		os.Exit(1)
	}
}

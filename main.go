package main

import (
	"fmt"
	"os"
	"os/exec"

	"git.sr.ht/~alphatroya/atr-capture/bookmarks"
	"git.sr.ht/~alphatroya/atr-capture/draft"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/forms"
	"git.sr.ht/~alphatroya/atr-capture/save"
	"github.com/charmbracelet/huh"
)

func main() {
	if _, err := env.CheckEnvs(); err != nil {
		fmt.Printf("Error in configuration: %s\n", err)
		os.Exit(1)
	}

	d := enterText()
	d, containURL, err := bookmarks.RequestTitleIfNeeded(d)
	checkErr("Form aborted: ", err, d)

	err = huh.NewConfirm().
		Title("Mark this note as TODO?").
		Value(&d.IsTODO).
		Run()
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

func enterText() draft.Draft {
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

	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	file, err := os.CreateTemp("", "*.md")
	if err != nil {
		fmt.Println("Error creating temp file: ", err)
		os.Exit(1)
	}
	file.Write([]byte(d.Text))
	cmd := exec.Command(editor, file.Name())

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error opening editor, editor=%s, path=%s, err=%v\n", editor, file.Name(), err)
		os.Exit(1)
	}
	r, err := os.ReadFile(file.Name())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file, path=%s, err=%v\n", file.Name(), err)
		os.Exit(1)
	}
	d.Text = string(r)
	defer os.Remove(file.Name())
	if d.Text == "" {
		fmt.Fprintf(os.Stderr, "File is empty, aborted, path=%s \n", file.Name())
		os.Exit(1)
	}
	return d
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

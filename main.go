package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"git.sr.ht/~alphatroya/atr-capture/bookmarks"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"git.sr.ht/~alphatroya/atr-capture/forms"
	"git.sr.ht/~alphatroya/atr-capture/save"
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

func main() {
	noteTitle := save.GenerateQuickNoteTitle(time.Now())
	err := save.SaveToJournal(noteTitle, envs.TodayJournalPath())
	checkErr("failed to add log to journal file: ", err)

	notePath := envs.PagePath(noteTitle)
	note, err := requestNoteFromUser(notePath)
	checkErr("failed to request note content from user: ", err)

	d, err := bookmarks.ExtractAndFormatLinkTitles(note)
	checkErr("failed to extract and format link titles from the note content: ", err)

	d.IsTODO, err = forms.RequestMarkAsTodo()
	checkErr("operation to request to mark the note as TODO was aborted or failed: ", err)

	saveContent := d.ContainURL() && forms.RequestSavingContent()
	err = save.SaveToPages(notePath, d, saveContent)
	checkErr("error writing the note to the file: ", err)

	fmt.Printf("Quick capture saved. A new note has been created: %s.md\n", noteTitle)
}

func requestNoteFromUser(path string) (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}

	cmd := exec.Command(editor, path)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("requestNoteFromUser: error opening editor, editor=%s, path=%s, err=%w\n", editor, path, err)
	}

	r, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("requestNoteFromUser: error reading file after saving, path=%s, err=%w\n", path, err)
	}

	text := string(r)
	if text == "" {
		os.Remove(path)
		return "", fmt.Errorf("requestNoteFromUser: note file is empty, aborted, file=%s \n", path)
	}
	return text, nil
}

func checkErr(message string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, message, err)
		os.Exit(1)
	}
}

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

	notePath := envs.PagePath(noteTitle)
	note := requestNoteFromUser(notePath)

	d, err := bookmarks.ExtractAndFormatLinkTitles(note)
	checkErr("page url title request failed: ", err)

	d.IsTODO, err = forms.RequestMarkAsTodo()
	checkErr("form aborted: ", err)

	saveContent := d.ContainURL() && forms.RequestSavingContent()
	err = save.SaveToPages(notePath, d, saveContent)
	checkErr("error writing to the file: ", err)

	fmt.Printf("quick capture saved, a new note created: %s.md\n", noteTitle)
}

func requestNoteFromUser(path string) string {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	cmd := exec.Command(editor, path)

	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error opening editor, editor=%s, path=%s, err=%v\n", editor, path, err)
		os.Exit(1)
	}

	r, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file, path=%s, err=%v\n", path, err)
		os.Exit(1)
	}

	text := string(r)
	if text == "" {
		fmt.Fprintf(os.Stderr, "file is empty, aborted, path=%s \n", path)
		os.Remove(path)
		os.Exit(1)
	}
	return text
}

func checkErr(message string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, message, err)
		os.Exit(1)
	}
}

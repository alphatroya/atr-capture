package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/alphatroya/atr-capture/bookmarks"
	"github.com/alphatroya/atr-capture/env"
	"github.com/alphatroya/atr-capture/forms"
	"github.com/alphatroya/atr-capture/save"
	"github.com/earthboundkid/versioninfo/v2"
)

var config env.Config

func init() {
	var err error
	var configDir string
	config, configDir, err = env.LoadConfig()
	if err != nil {
		fmt.Printf("error in configuration: can't load config file: %s\n", err)
		fmt.Printf(`
Create a similar config.json file at %s: 
{
	"path": ""
}`, configDir)
		fmt.Println("")
		os.Exit(1)
	}
}

func main() {
	versioninfo.AddFlag(nil)
	flag.Parse()

	noteTitle := save.GenerateQuickNoteTitle(time.Now())
	notePath := config.PagePath(noteTitle)

	var note string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		input, err := io.ReadAll(os.Stdin)
		checkErr("failed to read note content from stdin: ", err)
		err = os.WriteFile(notePath, input, 0644)
		checkErr("failed to write note content from stdin to given path: ", err)
		note = string(input)
	} else {
		var err error
		note, err = requestNoteFromUser(notePath)
		checkErr("failed to request note content from user: ", err)
	}

	err := save.SaveToJournal(noteTitle, config.TodayJournalPath())
	checkErr("failed to add log to journal file: ", err)

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
		return "", fmt.Errorf("requestNoteFromUser: error opening editor, editor=%s, path=%s, err=%w", editor, path, err)
	}

	r, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("requestNoteFromUser: error reading file after saving, path=%s, err=%w", path, err)
	}

	text := string(r)
	if text == "" {
		os.Remove(path)
		return "", fmt.Errorf("requestNoteFromUser: note file is empty, aborted, file=%s", path)
	}
	return text, nil
}

func checkErr(message string, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, message, err)
		os.Exit(1)
	}
}

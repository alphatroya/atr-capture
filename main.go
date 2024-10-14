package main

import (
	"errors"
	"fmt"
	"os"

	"git.sr.ht/~alphatroya/atr-capture/entry"
	"git.sr.ht/~alphatroya/atr-capture/env"
	"github.com/charmbracelet/huh"
)

func main() {
	err := env.CheckEnvs()
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	var text string
	huh.NewForm(
		huh.NewGroup(
			huh.NewText().
				Title("Capture this").
				ShowLineNumbers(true).
				Validate(func(in string) error {
					if len(in) == 0 {
						return errors.New("capture text can't be empty")
					}
					return nil
				}).
				Value(&text),
		),
	).
		Run()

	out := entry.NewEntry(text).Build()
	fmt.Println(out)
}

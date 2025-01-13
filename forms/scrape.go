package forms

import "github.com/charmbracelet/huh"

func RequestSavingContent() (confirm bool) {
	huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Save page content to ~/Downloads?").
				Value(&confirm),
		),
	).
		Run()
	return
}

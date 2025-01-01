package forms

import "github.com/charmbracelet/huh"

func RequestSavingContent() (confirm bool) {
	huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Request content?").
				Value(&confirm),
		),
	).
		Run()
	return
}

package forms

import (
	"github.com/charmbracelet/huh"
)

func ConfirmRestoreDraftDialog() (bool, error) {
	confirm := true
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Found a draft, use it?").
				Value(&confirm),
		),
	)

	err := form.Run()
	return confirm, err
}

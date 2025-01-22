package forms

import "github.com/charmbracelet/huh"

func RequestMarkAsTodo() (bool, error) {
	todo := false
	err := huh.NewConfirm().
		Title("Mark this note as TODO?").
		Value(&todo).
		Run()
	return todo, err
}

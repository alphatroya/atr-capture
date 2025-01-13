package forms

import (
	"git.sr.ht/~alphatroya/atr-capture/tags"
	"github.com/charmbracelet/huh"
)

func PickUpTags(in []tags.Item) (tags []string, err error) {
	options := make([]huh.Option[string], 0, len(tags))
	for _, t := range in {
		if t.Alias == "" {
			options = append(options, huh.NewOption(t.Name, t.Name))
			continue
		}
		options = append(options, huh.NewOption(t.Alias, t.Name))
	}

	err = huh.NewMultiSelect[string]().
		Title("Select tags").
		Options(options...).
		Value(&tags).
		Run()
	return
}

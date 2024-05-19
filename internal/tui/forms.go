package tui

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

func NewEntryForm() *huh.Form {

	var levels = huh.NewOptions[int](1, 2, 3, 4, 5, 6, 7)

	levels[0].Key = "1 - I anticipate that I would enjoy doing this."
	levels[1].Key = "2 - This seems like it would be relatively easy."
	levels[2].Key = "3 - This seems doable, although I'm not looking forward to it."
	levels[3].Key = "4 - It would take me some effort to do this."
	levels[4].Key = "5 - I'm somewhat anxious about trying this."
	levels[5].Key = "6 - I'm really daunted by this."
	levels[6].Key = "7 - I don't even want to think about this."

	var action string
	var thoughts string
	var startStrategy string
	done := true

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("What action are you considering?").Description("I am considering ...").
				Prompt("> ").
				Value(&action),
			huh.NewSelect[int]().
				Key("Level of difficulty").
				Options(levels...).
				Title("How hard does this seem, from 1 - 7?").
				Description("Choose one of the following"),
			huh.NewText().
				Title("What thoughts and feelings come to mind\nas you anticipate doing this?").Placeholder("write here ...").
				Value(&thoughts).CharLimit(-1),
			huh.NewText().
				Title("How would you start?").Placeholder("write here ...").
				Value(&startStrategy).CharLimit(-1),
			huh.NewConfirm().
				Key("done").
				Title("All done?").
				Validate(func(v bool) error {
					if !v {
						return fmt.Errorf("Take your time and reflect.")
					}
					return nil
				}).
				Affirmative("Yep").
				Negative("Wait, no").
				Value(&done),
		),
	).WithWidth(80).
		WithShowHelp(true).
		WithShowErrors(true)

	return form
}

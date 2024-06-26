package tui

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

const (
	KeyAction       string = "action"
	KeyDifficulty          = "difficulty"
	KeyThoughts            = "thoughts"
	KeyStrategy            = "strategy"
	KeyDone                = "done"
	KeyOutcomeValue        = "outcome"
	KeyReflection          = "reflection"
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
				Key(KeyAction).
				Title("What action are you considering?").Description("I am considering ...").
				Prompt("> ").
				Value(&action),
			huh.NewSelect[int]().
				Key(KeyDifficulty).
				Options(levels...).
				Title("How hard does this seem, from 1 - 7?").
				Description("Choose one of the following"),
			huh.NewText().
				Key(KeyThoughts).
				Title("What thoughts and feelings come to mind\nas you anticipate doing this?").Placeholder("write here ...").
				Value(&thoughts).CharLimit(-1),
			huh.NewText().
				Key(KeyStrategy).
				Title("How would you start?").Placeholder("write here ...").
				Value(&startStrategy).CharLimit(-1),
			huh.NewConfirm().
				Key(KeyDone).
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

func NewReflectionForm() *huh.Form {

	// TODO: need to do some thinking about how this value is actually
	// supposed to relate to the predicted difficulty of the action
	var levels = huh.NewOptions[int](1, 2, 3, 4, 5, 6, 7)
	levels[0].Key = "1 - That was really fun."
	levels[1].Key = "2 - ."
	levels[2].Key = "3 - ."
	levels[3].Key = "4 - ."
	levels[4].Key = "5 - ."
	levels[5].Key = "6 - ."
	levels[6].Key = "7 - That was horrible."

	var reflections string
	done := true

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Key(KeyOutcomeValue).
				Options(levels...).
				Title("How hard was this actually, from 1 - 7?").
				Description("Choose one of the following"),
			huh.NewText().
				Key(KeyReflection).
				Title("What was it like for you, doing this?").Placeholder("write here ...").
				Value(&reflections).CharLimit(-1),
			huh.NewConfirm().
				Key(KeyDone).
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

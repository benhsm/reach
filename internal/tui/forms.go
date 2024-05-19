package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 160

type state int

const (
	statusNormal state = iota
	stateDone
)

type EntryModel struct {
	state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewEntryModel() EntryModel {
	m := EntryModel{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

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

	m.form = huh.NewForm(
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
	).
		WithWidth(80).
		WithShowHelp(true).
		WithShowErrors(true)
	return m
}

func (m EntryModel) Init() tea.Cmd {
	return m.form.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m EntryModel) View() string {
	s := m.styles

	// Form (left side)
	v := strings.TrimSuffix(m.form.View(), "\n\n")
	form := m.lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		Render(v)

	// Status (right side)
	var status string
	{
		const statusWidth = 80
		status = s.Status.
			Height(lipgloss.Height(v)).
			Width(statusWidth).
			Render(s.StatusHeader.Render("Current Build"))
	}

	_ = lipgloss.JoinHorizontal(lipgloss.Top, form, status)
	return s.Base.Render(form)
}

func (m EntryModel) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

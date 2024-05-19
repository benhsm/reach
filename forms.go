package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 160

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

type state int

const (
	statusNormal state = iota
	stateDone
)

type Model struct {
	state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewAddActionModel() Model {
	m := Model{width: maxWidth}
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

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Model) View() string {
	s := m.styles

	// Form (left side)
	// v := strings.TrimSuffix(m.form.View(), "\n\n")
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

func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

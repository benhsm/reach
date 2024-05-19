package tui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

const maxWidth = 160

type state int

const (
	tableView state = iota
	entryView
)

type Model struct {
	state     state
	lg        *lipgloss.Renderer
	styles    *Styles
	entryForm *huh.Form
	width     int
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.entryForm = NewEntryForm()
	return m
}

func (m Model) Init() tea.Cmd {
	return m.entryForm.Init()
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if v, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = min(v.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	}

	var cmds []tea.Cmd

	switch m.state {
	case entryView:
		// Process the form
		form, cmd := m.entryForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.entryForm = f
			cmds = append(cmds, cmd)
		}
		if m.entryForm.State == huh.StateCompleted {
			// TODO: do something with the submitted data
			m.state = tableView
		}
		if m.entryForm.State == huh.StateAborted {
			m.state = tableView
		}
	default:
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "ctrl+c", "q":
				return m, tea.Quit
			case "a":
				m.entryForm = NewEntryForm()
				m.state = entryView
				cmds = append(cmds, m.entryForm.Init())
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	switch m.state {
	case entryView:
		v := strings.TrimSuffix(m.entryForm.View(), "\n\n")
		form := m.lg.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(indigo).
			Render(v)

		return s.Base.Render(form)
	default:
		return "you're looking at what's going to be the table view. Hit 'a' to add an action"
	}
}

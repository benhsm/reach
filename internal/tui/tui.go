package tui

import (
	"fmt"
	"strings"

	"github.com/benhsm/reach/internal/action"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"
)

const maxWidth = 160

type state int

const (
	tableFocus state = iota
	entryFocus
	reflectionFocus
)

type Model struct {
	state          state
	lg             *lipgloss.Renderer
	styles         *Styles
	entryForm      *huh.Form
	reflectionForm *huh.Form
	width          int

	table   table.Model
	actions []action.Action
}

func NewModel() Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	m.actions = []action.Action{
		{
			ID:            1,
			Desc:          "Example action",
			Difficulty:    4,
			Notes:         "Neque porro quisquam est qui dolorem ipsum quia dolor sit amet, consectetur, adipisci velit",
			Status:        action.StatusPending,
			StartStrategy: "Sis dos amet",
		},
		{
			ID:            2,
			Desc:          "Another action",
			Difficulty:    6,
			Notes:         "I'm scared of doing this because of X reason",
			Status:        action.StatusDone,
			StartStrategy: "Do the first thing",
			OutcomeValue:  4,
			Reflection:    "That wasn't as bad as I thought it would be",
		},
	}

	m.entryForm = NewEntryForm()
	m.table = NewTable(m.actions)
	return m
}

func (m Model) Init() tea.Cmd {
	return m.entryForm.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if v, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = min(v.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	}

	var cmds []tea.Cmd

	switch m.state {
	case entryFocus:
		// Process the form
		form, cmd := m.entryForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.entryForm = f
			cmds = append(cmds, cmd)
		}
		if m.entryForm.State == huh.StateCompleted {
			var lastActionId int64
			l := len(m.actions)
			if l != 0 {
				lastActionId = m.actions[l-1].ID
			}
			a := action.Action{
				ID:            lastActionId + 1,
				Status:        action.StatusPending,
				Notes:         m.entryForm.GetString(KeyThoughts),
				StartStrategy: m.entryForm.GetString(KeyStrategy),
				Difficulty:    m.entryForm.GetInt(KeyDifficulty),
				Desc:          m.entryForm.GetString(KeyAction),
			}
			m.actions = append(m.actions, a)
			m.table = NewTable(m.actions)
			m.state = tableFocus
		}
		if m.entryForm.State == huh.StateAborted {
			m.state = tableFocus
		}
	case reflectionFocus:
		// Process the form
		form, cmd := m.reflectionForm.Update(msg)
		if f, ok := form.(*huh.Form); ok {
			m.entryForm = f
			cmds = append(cmds, cmd)
		}
		if m.reflectionForm.State == huh.StateCompleted {
			idx := m.table.GetHighlightedRowIndex()
			m.actions[idx].OutcomeValue = m.reflectionForm.GetInt(KeyOutcomeValue)
			m.actions[idx].Reflection = m.reflectionForm.GetString(KeyReflection)
			m.actions[idx].Status = action.StatusDone
			m.table = NewTable(m.actions)
			m.state = tableFocus
		}
		if m.reflectionForm.State == huh.StateAborted {
			m.state = tableFocus
		}
	default: // table view
		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.String() {
			case "esc", "ctrl+c", "q":
				return m, tea.Quit
			case "a":
				m.entryForm = NewEntryForm()
				m.state = entryFocus
				cmds = append(cmds, m.entryForm.Init())
			case "r":
				m.reflectionForm = NewReflectionForm()
				m.state = reflectionFocus
				cmds = append(cmds, m.reflectionForm.Init())
			default:
				table, cmd := m.table.Update(msg)
				m.table = table
				cmds = append(cmds, cmd)
			}
		}
	}
	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	switch m.state {
	case entryFocus:
		v := strings.TrimSuffix(m.entryForm.View(), "\n\n")
		form := m.lg.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(indigo).
			Render(v)

		return lipgloss.JoinVertical(lipgloss.Center, s.Base.Render(form), "ctrl-c to cancel form entry and go back.")
	case reflectionFocus:
		v := strings.TrimSuffix(m.reflectionForm.View(), "\n\n")
		form := m.lg.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(indigo).
			Render(v)

		left := lipgloss.JoinVertical(lipgloss.Center, form, "ctrl-c to cancel form entry and go back.")

		idx := m.table.GetHighlightedRowIndex()
		selectedAction := m.actions[idx]
		right := m.lg.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(indigo).
			Width(40).
			Height(lipgloss.Height(left)).
			Render(displayAction(selectedAction))
		return s.Base.Render(lipgloss.JoinHorizontal(lipgloss.Top, left, right))
	default:
		help := `Help: 'a' to add a new possible action.
jk/↑↓ to change selection.
'r' to reflect on a the highlighted action.
'q' to quit`
		return s.Base.Render(lipgloss.JoinVertical(lipgloss.Bottom, m.table.View(), help))
	}
}

func displayAction(a action.Action) string {
	return fmt.Sprintf(
		`Action: %s
Predicted Difficulty: %d
Notes: %s
`, a.Desc, a.Difficulty, a.Notes)
}

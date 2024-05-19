package tui

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	p *tea.Program
)

type sessionState int

const (
	tableView sessionState = iota
	entryView
	reflectView
)

type MainModel struct {
	state     sessionState
	entryForm tea.Model
	// table       tea.Model
	// reflectForm tea.Model
}

func NewMainModel() MainModel {
	return MainModel{
		state:     entryView,
		entryForm: NewEntryModel(),
	}
}

func (m MainModel) Init() tea.Cmd {
	return m.entryForm.Init()
}

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c", "q":
			return m, tea.Quit
		default:
		}
	}

	switch m.state {
	case entryView:
		m.entryForm, cmd = m.entryForm.Update(msg)
		// case tableView:
		// 	m.table, cmd = m.table.Update(msg)
		// case reflectView:
		// 	m.table, cmd = m.reflectForm.Update(msg)
	}

	return m, tea.Batch(cmd)
}

func (m MainModel) View() string {
	switch m.state {
	case entryView:
		return m.entryForm.View()
	// case tableView:
	// 	return m.table.View()
	// case reflectView:
	// 	return m.reflectForm.View()
	default:
		return m.entryForm.View()
	}
}

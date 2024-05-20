package tui

import (
	"github.com/benhsm/reach/internal/action"
	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyDesc         = "description"
	columnKeyDifficulty   = "difficulty level"
	columnKeyNotes        = "notes"
	columnKeyStart        = "starting step"
	columnKeyReflection   = "reflection"
	columnKeyOutcomeValue = "outcome"
	columnKeyStatus       = "status"
)

func NewTable(actions []action.Action) table.Model {
	columns := []table.Column{
		// TODO: all the column widths here are guesses that need adjustment
		table.NewColumn(columnKeyDesc, columnKeyDesc, 20),
		table.NewColumn(columnKeyDifficulty, "difficulty", 10),
		table.NewColumn(columnKeyNotes, columnKeyNotes, 20),
		table.NewColumn(columnKeyStart, columnKeyStart, 20),
		table.NewColumn(columnKeyOutcomeValue, columnKeyOutcomeValue, 10),
		table.NewColumn(columnKeyReflection, columnKeyReflection, 20),
		table.NewColumn(columnKeyStatus, columnKeyStatus, 10),
	}

	rows := []table.Row{}
	for _, action := range actions {
		newRow := table.NewRow(
			table.RowData{
				columnKeyDesc:         action.Desc,
				columnKeyDifficulty:   action.Difficulty,
				columnKeyNotes:        action.Notes,
				columnKeyStart:        action.StartStrategy,
				columnKeyReflection:   action.Reflection,
				columnKeyOutcomeValue: action.OutcomeValue,
				columnKeyStatus:       action.Status,
			})
		rows = append(rows, newRow)
	}

	tableModel := table.New(columns).
		WithRows(rows).
		Focused(true).
		WithMultiline(true)

	return tableModel
}

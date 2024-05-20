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
		table.NewFlexColumn(columnKeyDesc, columnKeyDesc, 2),
		table.NewFlexColumn(columnKeyDifficulty, "difficulty", 1),
		table.NewFlexColumn(columnKeyNotes, columnKeyNotes, 4),
		table.NewFlexColumn(columnKeyStart, columnKeyStart, 3),
		table.NewFlexColumn(columnKeyOutcomeValue, columnKeyOutcomeValue, 1),
		table.NewFlexColumn(columnKeyReflection, columnKeyReflection, 3),
		table.NewFlexColumn(columnKeyStatus, columnKeyStatus, 1),
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
		WithMultiline(true).
		WithTargetWidth(100)

	return tableModel
}

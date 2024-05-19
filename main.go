package main

import (
	"fmt"
	"os"

	"github.com/benhsm/reach/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	_, err := tea.NewProgram(tui.NewEntryModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}

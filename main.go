package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benhsm/reach/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if f, err := tea.LogToFile("debug.log", "debug"); err != nil {
		fmt.Println("Couldn't open a file for logging:", err)
		os.Exit(1)
	} else {
		defer func() {
			err = f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	_, err := tea.NewProgram(tui.NewMainModel()).Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}
}

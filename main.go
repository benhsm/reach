package main

import (
	"database/sql"
	"log"

	constants "github.com/benhsm/reach/internal/constants"
	"github.com/benhsm/reach/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := initDb()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	model, err := tui.NewModel(db)
	if err != nil {
		log.Fatal(err)
	}
	_, err = tea.NewProgram(model).Run()
	if err != nil {
		log.Fatal(err)
	}
}

func initDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", constants.DatabasePath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(constants.SqlPrep); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

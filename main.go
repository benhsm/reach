package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/adrg/xdg"
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
	databasePath, err := xdg.DataFile("reach/reach.db")
	if err != nil {
		return nil, err
	}
	if _, err := os.Stat(databasePath); err != nil {
		fmt.Printf("reach requires an sqlite database to function. Create one at %s? (y/n): ", databasePath)
		in := bufio.NewReader(os.Stdin)
		answer, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		switch answer {
		case "y\n", "Y\n":
			fmt.Println("creating database.")
		default:
			fmt.Println("recieved non-affirmative answer. quitting.")
			os.Exit(0)
		}
	}
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec(constants.SqlPrep); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

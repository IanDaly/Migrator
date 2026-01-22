package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/iandaly/migrator/app"
	"github.com/iandaly/migrator/config"
	"github.com/iandaly/migrator/help"
)

func header() {
	header := `
 ██████   ██████  ███                                █████                      
░░██████ ██████  ░░░                                ░░███                       
 ░███░█████░███  ████   ███████ ████████   ██████   ███████    ██████  ████████ 
 ░███░░███ ░███ ░░███  ███░░███░░███░░███ ░░░░░███ ░░░███░    ███░░███░░███░░███
 ░███ ░░░  ░███  ░███ ░███ ░███ ░███ ░░░   ███████   ░███    ░███ ░███ ░███ ░░░ 
 ░███      ░███  ░███ ░███ ░███ ░███      ███░░███   ░███ ███░███ ░███ ░███     
 █████     █████ █████░░███████ █████    ░░████████  ░░█████ ░░██████  █████    
░░░░░     ░░░░░ ░░░░░  ░░░░░███░░░░░      ░░░░░░░░    ░░░░░   ░░░░░░  ░░░░░     
                       ███ ░███                                                 
                      ░░██████                                                  
                       ░░░░░░`

	fmt.Println(header)
}

func main() {

	if len(os.Args) == 1 {
		fmt.Println("Not enough arguments supplied")
		return
	}

	header()

	first := os.Args[1]

	// if we are initializing a project or 
	// just want to display the help menu
	switch first {
		case "init":
			exists, err := config.Exists()
			if err != nil {
				panic(err)
			}

			if exists {
				panic(errors.New("Config exists"))
			}

			if err := config.Create(); err != nil {
				panic(err)
			}

			fmt.Println("Initialize complete")
			return
		case "help":
				help.DisplayHelpMenu()
				return
	}

	// since we are not initializing the project
	// we need to create an app instance
	a := app.New()

	defer a.Driver.Close()

	if err := a.EnsureDriver(); err != nil {
		panic(err)
	}

	switch first {
	// if we are creating a migration
	case "make":
		if err := a.CreateMigration(); err != nil {
			panic(err)
		}

	case "migrate":
		if err := a.RunMigrations(); err != nil {
			panic(err)
		}

	// prints a list of pending migrations
	case "pending":
		if err := a.ListPendingMigrations(); err != nil {
			panic(err)
		}

	case "rollback":
		if err := a.Rollback(); err != nil {
			panic(err)
		}
	}
}

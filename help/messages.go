package help

import "fmt"

func DisplayHelpMenu() {

	helpMenu := `Usage:
	migrator init 				Initialize the tool in the current directory
		
	migrator migrate			Runs pending migrations 
		
	migrator pending			Output a list of migrations to be run
		
	migrator rollback [<steps>] 		Rollback applied migrations in the database by the number of steps, default 0 
	`

	fmt.Println(helpMenu)
}
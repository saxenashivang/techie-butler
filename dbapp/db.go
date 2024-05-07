package main

import (
	"net/http"

	"github.com/saxenashivang/techiebutler/config"
	"github.com/saxenashivang/techiebutler/database/commands"
	"github.com/spf13/cobra"
)

// dbapp runs as a cli tool to seed, migrate and drop tables
func main() {
	db()
}

func db() {
	config.LoadConfigs()
	cmd := &cobra.Command{}

	cmd.AddCommand(commands.DropTables())
	cmd.AddCommand(commands.Migrate())
	cmd.AddCommand(commands.Seed())
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	if config.App.Env != "development" {
		err = http.ListenAndServe(":8000", nil)
		if err != nil {
			panic(err.Error() + ". Error in DBMain")
		}
	}
}

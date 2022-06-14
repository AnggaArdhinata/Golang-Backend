package command

import (
	"backend-golang/src/configs/serve"
	database "backend-golang/src/database/gorm"

	"github.com/spf13/cobra"
)

var initCommand = cobra.Command{
	Use: "Backend Golang",
	Short: "Simple API With Golang",
}

func init() {
	initCommand.AddCommand(serve.ServeCmd)
	initCommand.AddCommand(database.MigrateCmd)
}

func Run(args []string) error {
	initCommand.SetArgs(args)

	return initCommand.Execute()
}
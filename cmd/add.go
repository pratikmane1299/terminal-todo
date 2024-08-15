package cmd

import (
	"fmt"

	"github.com/pratikmane1299/terminal-todo/db"
	"github.com/pratikmane1299/terminal-todo/utils"
	"github.com/spf13/cobra"
)

var addTodoCmd = &cobra.Command{
	Use:   "add [OPTIONS]",
	Short: "Add a new todo",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		todo := args[0]

		todoDb, err := db.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}

		defer todoDb.Db.Close()

		err = todoDb.AddTodo(todo)
		if err != nil {
			return fmt.Errorf("unable to add new todos :%w", err)
		}

		fmt.Printf("%s todo added successfully", todo)

		return nil
	},
}

package cmd

import (
	"fmt"
	"strconv"

	"github.com/pratikmane1299/terminal-todo/db"
	"github.com/pratikmane1299/terminal-todo/utils"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [OPTIONS]",
	Short: "Delete a todo by it's id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		todoDb, err := db.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}

		todoId, _ := strconv.Atoi(args[0])

		err = todoDb.DeleteTodoById(todoId)
		if err != nil {
			return fmt.Errorf("could not delete todo by id %w", err)
		}

		fmt.Printf("todo with id %d deleted successfully", todoId)

		return nil
	},
}

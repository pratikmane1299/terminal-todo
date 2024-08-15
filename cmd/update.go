package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pratikmane1299/terminal-todo/db"
	"github.com/pratikmane1299/terminal-todo/utils"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [OPTIONS] [FLAGS]",
	Short: "Update a todo by its id",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		todoId, _ := strconv.Atoi(args[0])

		todo, err := cmd.Flags().GetString("todo")
		if err != nil {
			return fmt.Errorf("--todo flag is required")
		}

		completed, err := cmd.Flags().GetBool("completed")
		if err != nil {
			return fmt.Errorf("--completed flag is required")
		}

		todoDb, err := db.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}

		err = todoDb.UpdateTodo(db.Todo{
			Id:        todoId,
			Todo:      todo,
			Completed: completed,
			Created:   time.Time{},
		})
		if err != nil {
			return fmt.Errorf("could not update todo %w", err)
		}

		return nil
	},
}

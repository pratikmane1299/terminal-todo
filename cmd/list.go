package cmd

import (
	"fmt"

	"github.com/pratikmane1299/terminal-todo/db"
	"github.com/pratikmane1299/terminal-todo/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all your todos",
	RunE: func(cmd *cobra.Command, args []string) error {
		todoDb, err := db.OpenDB(utils.SetupPath())
		if err != nil {
			return err
		}

		defer todoDb.Db.Close()

		status, err := cmd.Flags().GetString("status")
		if err != nil {
			return nil
		}

		todos, err := todoDb.GetAllTodos(status)
		if err != nil {
			return fmt.Errorf("unable to fetch todos :%w", err)
		}

		var output = ""

		fmt.Println("")

		if len(todos) > 0 {
			for _, todo := range todos {
				var completed = " "

				if todo.Completed {
					completed = "x"
				}

				output += fmt.Sprintf("[%s] %d. %s \n\n", completed, todo.Id, todo.Todo)
			}
		} else {
			output = "You got no todos homie, how about you add some huhhhh....\n"
		}

		fmt.Println(output)

		return nil
	},
}

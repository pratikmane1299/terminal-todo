package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "terminal-todo",
	Short: "A Cli based todo app",
	Long:  "Teminal Todo is an terminal based todo app for you nerds out there.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return cmd.Help()
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	listCmd.Flags().StringP("status", "s", "", "filter todos by status completed/pending/all")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(addTodoCmd)
	rootCmd.AddCommand(deleteCmd)

	updateCmd.Flags().StringP("todo", "t", "", "todo to update")
	updateCmd.Flags().BoolP("completed", "c", false, "mark todo as complete/pending")

	rootCmd.AddCommand(updateCmd)
}

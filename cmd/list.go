package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks in your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("List is called")
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

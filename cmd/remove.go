package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the task from your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		for _, t := range args {
			task_no, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("Could not understand the arg:", t)
				continue
			}
			fmt.Println("Task number", task_no, "is removed from the list")
		}
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

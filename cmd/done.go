package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Marks the task as done/completed",
	Run: func(cmd *cobra.Command, args []string) {

		for _, t := range args {
			task_no, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("Could not understand the arg:", t)
				continue
			}
			fmt.Println("Task number", task_no, "is marked as Done")
		}
	},
}

func init() {
	RootCmd.AddCommand(doneCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strings"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if task == "" {
			fmt.Println("Nothing to add from your input")
			return
		}
		task_id, err := todo_db.AddTask(task)
		_ = task_id
		if err != nil {
			fmt.Println("Something went wrong :(", err.Error())
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your to-do list.\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

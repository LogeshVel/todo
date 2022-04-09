package cmd

import (
	"fmt"
	"os"
	"strconv"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove the task from your to-do list.",
	Example: `
	todo rm 1 2
		- here 1 and 2 are the Task number of the active task
	todo rm 3
		- here 3 is the Task number of the active task
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks_map, err := todo_db.Get_active_task()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		for _, t := range args {
			task_no, err := strconv.Atoi(t)
			if err != nil {
				fmt.Println("Could not understand the arg:", t)
				continue
			}
			if task_value, found := tasks_map[task_no]; found {
				d_err := todo_db.Remove_task(task_value.Key, task_value.Value)
				if d_err != nil {
					fmt.Println("Something went wrong :( while removing the Task number :", task_no, d_err.Error())
					continue

				}
				fmt.Printf("Task \"%d. %s\" is removed from your To-Do list\n", task_no, task_value.Value)
				// fmt.Println("Task number", task_no, "is removed from the list")
				continue
			} else {
				fmt.Println("The Task number", task_no, "does not exists")
			}

		}
		Display_all_active_task()
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

package cmd

import (
	"fmt"
	"os"
	"strconv"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "del",
	Short: "Permanently remove/delete the task from your task list.",
	Example: `
	todo del 1 2
		- here 1 and 2 are the Task number of the active task"
	todo del 3
		- here 3 is the Task number of the active task"
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
				d_err := todo_db.DeleteTask(task_value.Key)
				if d_err != nil {
					fmt.Println("Something went wrong :( while deleting the Task number :", task_no, d_err.Error())
					continue

				}
				fmt.Printf("Task \"%d. %s\" is permanently deleted from your To-Do list\n", task_no, task_value.Value)
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
	RootCmd.AddCommand(deleteCmd)
}

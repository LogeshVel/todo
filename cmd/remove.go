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
	Short: "Remove the task from your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		tasks_map, err := todo_db.AllTasks()
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
				fmt.Printf("Task \"%d. %s\" is removed from your To-Do list\n", task_no, task_value.Value)
				// fmt.Println("Task number", task_no, "is removed from the list")
				continue
			} else {
				fmt.Println("The Task number", task_no, "does not exists")
			}

		}
		Display_all_task()
	},
}

func init() {
	RootCmd.AddCommand(removeCmd)
}

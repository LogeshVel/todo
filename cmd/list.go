package cmd

import (
	"fmt"
	"os"
	"sort"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks in your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		Display_all_task()
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}

func Display_all_task() {
	tasks, err := todo_db.AllTasks()
	if err != nil {
		fmt.Println("Something went wrong while listing all tasks:", err)
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("\nYou have no tasks to complete! Try adding new tasks")
		return
	}
	keys := make([]int, 0)
	for k := range tasks {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("\nYou have the following tasks in your to-do list:")
	for _, sorted_key := range keys {
		fmt.Printf("%d. %s\n", sorted_key, tasks[sorted_key].Value)
	}
}

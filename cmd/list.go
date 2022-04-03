package cmd

import (
	"fmt"
	"os"
	"sort"
	"todo/colors"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the tasks in your task list.",
	Example: `
	todo list
	todo list -a
		- both cmds lists all the active tasks in your to-do list
	todo list -r
		- shows the tasks that have been removed from the to-do list
	todo list -d
		- shows the tasks that have been marked as done
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			a_status bool
			r_status bool
			d_status bool
		)
		if a_status, _ = cmd.Flags().GetBool("active"); a_status {
			Display_all_active_task()
		}
		if r_status, _ = cmd.Flags().GetBool("remove"); r_status {
			Display_all_removed_task()
		}
		if d_status, _ = cmd.Flags().GetBool("done"); d_status {
			Display_all_done_task()
		}

		if !(a_status || r_status || d_status) {
			Display_all_active_task()
		}

	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
	ListCmd.Flags().BoolP("active", "a", false, "Lists active to-do Tasks")
	ListCmd.Flags().BoolP("remove", "r", false, "Lists removed to-do Tasks")
	ListCmd.Flags().BoolP("done", "d", false, "Lists to-do tasks that are marked as Done")

}

func Display_all_active_task() {
	tasks, err := todo_db.Get_active_task()
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
	fmt.Printf(string(colors.ColorPurple))
	fmt.Println("Task No. Task", string(colors.ColorReset))
	for _, sorted_key := range keys {
		fmt.Printf("      %d. %s\n", sorted_key, tasks[sorted_key].Value)
	}
}

func Display_all_removed_task() {
	tasks, err := todo_db.Get_removed_task()
	if err != nil {
		fmt.Println("Something went wrong while listing removed tasks:", err)
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("\nYou have no tasks that have removed")
		return
	}
	keys := make([]int, 0)
	for k := range tasks {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("\nYou have removed the following tasks from your to-do list:")
	fmt.Printf(string(colors.ColorRed))
	fmt.Println("Task No. Task", string(colors.ColorReset))
	for _, sorted_key := range keys {
		fmt.Printf("      %d. %s\n", sorted_key, tasks[sorted_key].Value)
	}
}

func Display_all_done_task() {
	tasks, err := todo_db.Get_done_task()
	if err != nil {
		fmt.Println("Something went wrong while listing done tasks:", err)
		os.Exit(1)
	}
	if len(tasks) == 0 {
		fmt.Println("\nYou have no tasks that have marked as done!")
		return
	}
	keys := make([]int, 0)
	for k := range tasks {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	fmt.Println("\nYou have Done the following tasks in your to-do list:")
	fmt.Printf(string(colors.ColorGreen))
	fmt.Println("Task No. Task", string(colors.ColorReset))
	for _, sorted_key := range keys {
		fmt.Printf("      %d. %s\n", sorted_key, tasks[sorted_key].Value)
	}
}

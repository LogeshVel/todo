package cmd

import (
	"fmt"
	"os"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var clearrmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Clears the removed tasks list",
	Example: `
	todo clear rm
		- clears all the removed tasks list
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks_map, err := todo_db.Get_removed_task()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Print("Clearning removed tasks list... ")
		if len(tasks_map) == 0 {
			fmt.Print("Removed tasks list is Empty!\n")
			return
		}
		for _, task := range tasks_map {
			d_err := todo_db.DeleteTask(task.Key)
			if d_err != nil {
				fmt.Println("\n[ ✘ ]Something went wrong :( while clearning removed task lists :", d_err)

				continue

			}
			fmt.Print("✔")
		}
		fmt.Print(" cleared\n")
	},
}

func init() {
	clearCmd.AddCommand(clearrmCmd)

}

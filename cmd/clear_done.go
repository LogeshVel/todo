package cmd

import (
	"fmt"
	"os"
	todo_db "todo/db"

	"github.com/spf13/cobra"
)

var cleardoneCmd = &cobra.Command{
	Use:   "done",
	Short: "Clears the done tasks list",
	Example: `
	todo clear done
		- clears all the done tasks list
	`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks_map, err := todo_db.Get_done_task()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Print("Clearning done tasks list... ")
		if len(tasks_map) == 0 {
			fmt.Print("Done tasks list is Empty!\n")
			return
		}
		for _, task := range tasks_map {
			d_err := todo_db.DeleteTask(task.Key)
			if d_err != nil {
				fmt.Println("\n[ ✘ ]Something went wrong :( while clearning done tasks list :", d_err)

				continue

			}
			fmt.Print("✔")
		}
		fmt.Print(" cleared\n")
	},
}

func init() {
	clearCmd.AddCommand(cleardoneCmd)

}

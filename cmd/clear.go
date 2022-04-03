package cmd

import (
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Clears the done/removed list",
	Example: `
	todo clear done
		- clears the done lists
	todo clear rm
		- clears the removed task list
	`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	RootCmd.AddCommand(clearCmd)

}

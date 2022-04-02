package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "todo is a CLI To-Do list manager",
}

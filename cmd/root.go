package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: "tasker",
	Short: "A task list for your basic activities.",
}
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task into your task list.",
	Run: func(cmd *cobra.Command, args []string){
		fmt.Println("adds called")
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}
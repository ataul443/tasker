package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var addCmd = &cobra.Command{
	Use: "add",
	Short: "Adds a task into your task list.",
	Run: func(cmd *cobra.Command, args []string){
		taskDesc := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" into your task list.\n", taskDesc)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}
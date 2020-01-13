package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark tasks completed from your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int

		for _, idx := range args {
			id, err := strconv.Atoi(idx)
			if err != nil {
				fmt.Printf("Failed to parse \"%s\"\n", idx)
			}else {
				ids = append(ids, id)
			}
		}
		fmt.Printf("%v\n", ids)
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

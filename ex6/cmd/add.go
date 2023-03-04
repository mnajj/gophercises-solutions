package cmd

import (
	"ex6/db"
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

func init() {
	RootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add task to your todo list",
	Run: func(cmd *cobra.Command, args []string) {
		in := strings.Join(args, " ")
		if _, err := db.CreateTask(in); err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", in)
	},
}

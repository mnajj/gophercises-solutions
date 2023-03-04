package cmd

import (
	"ex6/db"
	"fmt"
	"github.com/spf13/cobra"
	"log"
)

func init() {
	RootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all your todo tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			log.Fatal("Something went wrong:", err)
		}
		if len(tasks) == 0 {
			fmt.Println("You don't have tasks today!")
			return
		}
		for i, t := range tasks {
			fmt.Printf("%d. %s\n", i+1, t.Value)
		}
	},
}

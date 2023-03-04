package cmd

import (
	"ex6/db"
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func init() {
	RootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "do task in your list by insert it's number",
	Run:   run,
}

func run(_ *cobra.Command, args []string) {
	ids := make([]int, 0, len(args))
	for _, i := range args {
		id, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("unable to parse \"%s\"\n", i)
			continue
		}
		ids = append(ids, id)
	}
	tasks, err := db.AllTasks()
	if err != nil {
		fmt.Println("Something went wrong:", err)
		return
	}
	for _, id := range ids {
		if id <= 0 || id > len(tasks) {
			fmt.Println("Invalid task number:", id)
			continue
		}
		task := tasks[id-1]
		if err := db.DeleteTask(task.Key); err != nil {
			fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err)
			return
		}
		fmt.Printf("Marked \"%d\" as completed.\n", id)
	}
}

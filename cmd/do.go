package cmd

import (
	"fmt"
	"go-excercises/task/db"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a task as complete",
	Run: func(cmd *cobra.Command, args []string) {
		var ids = []int{}
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Could not parse the arugument:", arg)
			} else {
				ids = append(ids, id)
			}
		}

		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			err := db.DeleteTask(tasks[id-1].Key)
			if err != nil {
				fmt.Printf("Failed to mark task \"%d\" as complete: %s\n", id, err)
			} else if err := db.StoreCompleteTask(tasks[id-1].Value); err != nil {
				fmt.Printf("Failed to mark task \"%d\" as complete: %s\n", id, err)
			} else {
				fmt.Printf("Task \"%d\" marked as complete\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}

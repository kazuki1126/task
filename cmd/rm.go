package cmd

import (
	"fmt"
	"go-excercises/task/db"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete task(s) from the task list",
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
				fmt.Printf("Failed to delete the task %d: %s\n", id, err)
			} else {
				fmt.Printf("Deleted the task %s from task list\n", tasks[id-1].Value)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(rmCmd)
}

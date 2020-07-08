package cmd

import (
	"fmt"
	"go-excercises/task/db"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a task to your task list",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("task added: \"%s\"\n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}

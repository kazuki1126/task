package cmd

import (
	"fmt"
	"go-excercises/task/db"
	"os"

	"github.com/spf13/cobra"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "Lists all completed tasks within the last 24 hours",
	Run: func(cmd *cobra.Command, _ []string) {
		doneTasks, err := db.AllCompleteTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(doneTasks) == 0 {
			fmt.Println(`No tasks marked as complete. Don't worry. Just chill out.`)
			return
		}
		fmt.Println("You have finished the following tasks today:")
		for _, doneTask := range doneTasks {
			fmt.Printf("- %s\n", doneTask)
		}
	},
}

func init() {
	RootCmd.AddCommand(completedCmd)
}

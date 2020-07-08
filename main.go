package main

import (
	"fmt"
	"go-excercises/task/cmd"
	"go-excercises/task/db"
	"os"
	"path/filepath"
)

func main() {
	home, _ := os.UserHomeDir()
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

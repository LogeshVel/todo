package main

import (
	"fmt"
	"os"
	"path/filepath"
	"todo/cmd"
	todo_db "todo/db"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbpath := filepath.Join(home, "tasks.db")
	shouldbe(todo_db.InitDB(dbpath))
	shouldbe(cmd.RootCmd.Execute())
}

func shouldbe(err error) {
	if err != nil {
		fmt.Println("✘", err.Error())
		os.Exit(1)
	}
}

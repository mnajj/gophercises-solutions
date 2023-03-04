package main

import (
	"ex6/cmd"
	"ex6/db"
	"github.com/mitchellh/go-homedir"
	"log"
	"path/filepath"
)

func main() {
	home, _ := homedir.Dir()
	path := filepath.Join(home, "tasks.db")
	must(db.Connect(path))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

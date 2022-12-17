package main

import (
	"fmt"
	"os"
	"tut/subcommand"
)

func main() {
	switch os.Args[1] {
	case "init":
		subcommand.Init(os.Args[2:])
	case "remote":
		subcommand.Remote(os.Args[2:])
	case "add":
		subcommand.Add(os.Args[2:])
	case "commit":
		subcommand.Commit(os.Args[2:])
	case "start-section":
		subcommand.StartSection(os.Args[2:])
	case "push":
		subcommand.Push(os.Args[2:])
	default:
		panic(fmt.Errorf("未知のサブコマンド"))
	}
}

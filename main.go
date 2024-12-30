package main

import (
	"fmt"
	"os"
	"myvcs/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: vcs <command> [options]")
		return
	}

	command := os.Args[1]
	switch command {
	case "start":
		cmd.InitRepo()
	case "keep":
		if len(os.Args) < 3 {
			fmt.Println("Usage: vcs keep <file>")
			return
		}
		cmd.AddFile(os.Args[2])
	case "save":
		if len(os.Args) < 3 {
			fmt.Println("Usage: vcs save <message>")
			return
		}
		cmd.SaveCommit(os.Args[2])
	case "timeline":
		cmd.ShowCommits()
	default:
		fmt.Println("Unknown command:", command)
	}
}

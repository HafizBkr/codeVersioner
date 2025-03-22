package main

import (
	"fmt"
	"os"
	"myvcs/cmd"
)

func main() {
	if len(os.Args) < 2 {
		cmd.ShowHelp()
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
	case "status":
		cmd.Status()
	case "branch":
		if len(os.Args) < 3 {
			fmt.Println("Usage: vcs branch --list | --create <name>")
			return
		}
		switch os.Args[2] {
		case "--list":
			cmd.ListBranches()
		case "--create":
			if len(os.Args) < 4 {
				fmt.Println("Usage: vcs branch --create <name>")
				return
			}
			cmd.CreateBranch(os.Args[3])
		default:
			fmt.Println("Usage: vcs branch --list | --create <name>")
		}
	case "clone":
		if len(os.Args) < 3 {
			fmt.Println("Usage: vcs clone <local_path>")
			return
		}
		cmd.CloneLocalRepo(os.Args[2])
	case "help":
		cmd.ShowHelp()
	default:
		fmt.Printf("Unknown command: %s\nUse 'vcs help' to see available commands\n", command)
	}
}

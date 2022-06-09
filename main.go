package main

import (
	"fmt"
	"os"

	"github.com/joe-reed/laminar/cli"
	"github.com/joe-reed/laminar/store"
)

func main() {
	store := store.FileStore{Path: "list.txt"}
	output := os.Stderr

	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) == 2 {
			fmt.Println("Paramter for \"add\" missing")
			printUsage()
			return
		}

		cli.Add(os.Args[2], store, output)
	case "next":
		cli.Next(store, output)
	case "done":
		cli.Done(store, output)
	case "help":
		printUsage()
	default:
		fmt.Println("Unrecognised command")
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("    ./bin/laminar add \"My new item\"")
	fmt.Println("    ./bin/laminar next")
	fmt.Println("    ./bin/laminar done")
}

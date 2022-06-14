package main

import (
	"fmt"
	"os"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/cli"
	"github.com/joe-reed/laminar/store"
)

func main() {
	store := store.FileStore{Path: "list.txt"}

	c := cli.Cli{
		Store:  store,
		Output: os.Stderr,
	}

	if len(os.Args) == 1 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) == 2 {
			fmt.Println("Parameter for \"add\" missing")
			printUsage()
			return
		}

		c.Add(os.Args[2])
	case "next":
		c.Next()
	case "done":
		c.Done()
	case "serve":
		api.Serve(store)
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

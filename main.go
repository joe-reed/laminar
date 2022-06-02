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
	switch os.Args[1] {
	case "add":
		cli.Add(os.Args[2], store, output)
	case "next":
		cli.Next(store, output)
	case "done":
		cli.Done(store, output)
	default:
		fmt.Printf("Unrecognised command. Usage:\n    ./bin/laminar add \"My new item\"\n    ./bin/laminar next\n")
	}
}

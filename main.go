package main

import (
	"os"

	"github.com/joe-reed/laminar/cli"
	"github.com/joe-reed/laminar/store"
)

func main() {
	cli.Add(os.Args[1], store.FileStore{Path: "list.txt"}, os.Stderr)
}

package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/cli"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func main() {
	cf := config.ConfigFile{Path: "config.txt"}
	config := cf.GetConfig()

	var s store.Store
	switch config.Store {
	case "api":
		s = store.ApiStore{BaseURL: config.Path, Client: http.DefaultClient}
	case "file":
		s = store.FileStore{Path: config.Path}
	}

	c := cli.Cli{
		Store:  s,
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
		api.Serve(s)
	case "configure":
		if len(os.Args) == 2 {
			fmt.Println("Parameter for \"configure\" missing")
			printUsage()
			return
		}

		cf.SetStore(os.Args[2])
		fmt.Printf("Configured: %s\n", cf.GetConfig().Path)
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
	fmt.Println("    ./bin/laminar serve")
	fmt.Println("    ./bin/laminar configure \"http://my-api-url.test")
}

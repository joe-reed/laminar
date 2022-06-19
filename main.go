package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func main() {
	cf := config.ConfigFile{Path: getConfigPath()}
	config := cf.GetConfig()

	var s store.Store
	switch config.Store {
	case "api":
		s = store.ApiStore{BaseURL: config.Path, Client: http.DefaultClient}
	case "file":
		s = store.FileStore{Path: config.Path}
	}

	cmd.Execute(s, cf)
}

func getConfigPath() string {
	dirname, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(dirname, ".laminar"), os.ModePerm)

	if err != nil {
		panic(err)
	}

	return filepath.Join(dirname, ".laminar", "config.txt")
}

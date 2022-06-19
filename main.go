package main

import (
	"os"
	"path/filepath"

	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
)

func main() {
	cf := config.ConfigFile{Path: getConfigPath()}

	cmd.Execute(cf.GetStore(), cf)
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

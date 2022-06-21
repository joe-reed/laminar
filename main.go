package main

import (
	"os"
	"path/filepath"

	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func main() {
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)

	c, err := config.Load(filepath.Join(home, ".laminar/config.yaml"))
	cobra.CheckErr(err)

	err = cmd.Execute(store.FromPath(c.GetStorePath()), c)
	cobra.CheckErr(err)
}

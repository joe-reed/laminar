package main

import (
	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func main() {
	c, err := config.Load()
	cobra.CheckErr(err)

	cmd.Execute(store.FromPath(c.GetStorePath()), c)
}

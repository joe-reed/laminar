package cmd

import (
	"fmt"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewConfigureCommand(s store.Store, cf config.ConfigFile) *cobra.Command {
	return &cobra.Command{
		Use:   "configure [file path or url for store]",
		Short: "Configure the store used by Laminar",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cf.SetStore(args[0])
			fmt.Printf("Configured: %s\n", cf.GetConfig().Path)
		},
	}
}

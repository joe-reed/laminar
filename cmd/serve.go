package cmd

import (
	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewServeCommand(s store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serve your Laminar instance as an API for use by the configure command",
		Run: func(cmd *cobra.Command, args []string) {
			api.Serve(s)
		},
	}
}

package cmd

import (
	"fmt"

	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewAddCommand(s store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "add [item to add]",
		Short: "Add an item to your list",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			s.Add(args[0])

			fmt.Fprintln(cmd.OutOrStdout(), "Item added")
		},
	}
}

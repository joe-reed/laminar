package cmd

import (
	"fmt"

	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewNextCommand(s store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "next",
		Short: "Get the next item from the list",
		Run: func(cmd *cobra.Command, args []string) {
			if s.Next() == "" {
				fmt.Fprintln(cmd.OutOrStdout(), "All items complete!")
				return
			}

			fmt.Fprintln(cmd.OutOrStdout(), s.Next())
		},
	}
}

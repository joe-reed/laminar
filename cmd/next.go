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
		RunE: func(cmd *cobra.Command, args []string) error {
			next, err := s.Next()
			if err != nil {
				return err
			}

			if next == "" {
				fmt.Fprintln(cmd.OutOrStdout(), "All items complete!")
				return nil
			}

			fmt.Fprintln(cmd.OutOrStdout(), next)
			return nil
		},
	}
}

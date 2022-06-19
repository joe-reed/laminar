package cmd

import (
	"fmt"

	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewDoneCommand(s store.Store) *cobra.Command {
	return &cobra.Command{
		Use:   "done",
		Short: "Complete the next item on your list",
		Run: func(cmd *cobra.Command, args []string) {
			s.Pop()
			fmt.Fprintln(cmd.OutOrStdout(), "Item complete")

			if s.Next() == "" {
				fmt.Fprintln(cmd.OutOrStdout(), "All items complete!")
				return
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Next: %s\n", s.Next())
		},
	}
}

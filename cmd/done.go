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
		RunE: func(cmd *cobra.Command, args []string) error {
			done, err := s.Pop()
			if err != nil {
				return err
			}

			if done != "" {
				fmt.Fprintf(cmd.OutOrStdout(), "Completed: %s\n", done)
			}

			next, err := s.Next()
			if err != nil {
				return err
			}

			if next == "" {
				fmt.Fprintln(cmd.OutOrStdout(), "All items complete!")
				return nil
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Next: %s\n", next)
			return nil
		},
	}
}

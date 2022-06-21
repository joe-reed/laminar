package cmd

import (
	"fmt"

	"github.com/joe-reed/laminar/config"
	"github.com/spf13/cobra"
)

func NewConfigureCommand(c *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "configure [file path or url for store]",
		Short: "Configure the store used by Laminar",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			c.SetStorePath(args[0])
			err := c.Write()
			if err != nil {
				return err
			}

			fmt.Fprintf(cmd.OutOrStdout(), "Configured: %s\n", args[0])

			return nil
		},
	}
}

package cmd

import (
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/spf13/cobra"
)

func NewRootCommand(s store.Store, c *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "laminar",
		Short: "CLI todo list for focus and flow.",
		Long: `A FIFO CLI todo list app to help keep you focussed on the next most important thing.
Work alone with a file-based list, or collaborate with others through simple built-in sharing features.`,
	}

	cmd.AddCommand(NewAddCommand(s))
	cmd.AddCommand(NewNextCommand(s))
	cmd.AddCommand(NewDoneCommand(s))
	cmd.AddCommand(NewConfigureCommand(c))
	cmd.AddCommand(NewServeCommand(s))

	return cmd
}

func Execute(s store.Store, c *config.Config) {
	err := NewRootCommand(s, c).Execute()
	cobra.CheckErr(err)
}

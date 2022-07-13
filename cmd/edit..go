package cmd

import (
	"os"
	"os/exec"

	"github.com/joe-reed/laminar/config"
	"github.com/spf13/cobra"
)

func NewEditCommand(c *config.Config) *cobra.Command {
	return &cobra.Command{
		Use:   "edit",
		Short: "Edit your list of items",
		RunE: func(cmd *cobra.Command, args []string) error {
			vim := exec.Command("vim", c.GetStorePath())
			vim.Stdin = os.Stdin
			vim.Stdout = os.Stdout
			return vim.Run()
		},
	}
}

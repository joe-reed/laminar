package cmd_test

import (
	"bytes"
	"testing"

	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func runCommand(t *testing.T, s store.Store, c *config.Config, args []string) string {
	cmd := cmd.NewRootCommand(s, c)

	output := &bytes.Buffer{}
	cmd.SetOut(output)

	cmd.SetArgs(args)

	err := cmd.Execute()

	if err != nil {
		t.Error(err)
	}

	return output.String()
}

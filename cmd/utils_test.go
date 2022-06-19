package cmd_test

import (
	"bytes"
	"testing"

	"github.com/joe-reed/laminar/cmd"
	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func runCommand(t *testing.T, s store.Store, args []string) string {
	cmd := cmd.NewRootCommand(s, config.ConfigFile{Path: t.TempDir()})

	output := &bytes.Buffer{}
	cmd.SetOut(output)

	cmd.SetArgs(args)

	cmd.Execute()

	return output.String()
}

package cmd_test

import (
	"os"
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func Test_it_sets_store_config(t *testing.T) {
	c, path := getConfig()
	defer os.Remove(path)

	runConfigure(t, c, "test.txt")

	got := c.GetStorePath()
	want := "test.txt"

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_success_message(t *testing.T) {
	v, path := getConfig()
	defer os.Remove(path)

	got := runConfigure(t, v, "test.txt")

	want := "Configured: test.txt\n"

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func getConfig() (c *config.Config, path string) {
	path = "./config_test.yaml"
	c, _ = config.Load(path)

	return
}

func runConfigure(t *testing.T, c *config.Config, path string) string {
	return runCommand(t, &store.InMemoryStore{}, c, []string{"configure", path})
}

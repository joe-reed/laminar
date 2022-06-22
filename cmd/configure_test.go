package cmd_test

import (
	"os"
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
)

func Test_it_sets_store_config(t *testing.T) {
	c, path := getConfig()
	defer os.Remove(path)

	runConfigure(t, c, "test.txt")

	assert.Equal(t, "test.txt", c.GetStorePath())
}

func Test_it_outputs_success_message(t *testing.T) {
	v, path := getConfig()
	defer os.Remove(path)

	assert.Equal(t, "Configured: test.txt\n", runConfigure(t, v, "test.txt"))
}

func getConfig() (c *config.Config, path string) {
	path = "./config_test.yaml"
	c, _ = config.Load(path)

	return
}

func runConfigure(t *testing.T, c *config.Config, path string) string {
	return runCommand(t, &store.InMemoryStore{}, c, []string{"configure", path})
}

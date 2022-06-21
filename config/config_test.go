package config_test

import (
	"os"
	"testing"

	"github.com/joe-reed/laminar/config"
)

func Test_it_sets_store_config(t *testing.T) {
	c := config.New()

	c.SetStorePath("/path/to/store.txt")

	got := c.GetStorePath()
	want := "/path/to/store.txt"

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func Test_it_loads_and_writes_to_file(t *testing.T) {
	path := "./config_test.yaml"
	c, _ := config.Load(path)

	c.SetStorePath("/path/to/store.txt")
	c.Write()

	newPath := "./config_test_new.yaml"
	defer os.Remove(newPath)

	os.Rename(path, newPath)

	c, _ = config.Load(newPath)

	got := c.GetStorePath()
	want := "/path/to/store.txt"

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

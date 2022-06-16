package config_test

import (
	"os"
	"testing"

	"github.com/joe-reed/laminar/config"
)

func Test_setting_api_store_config(t *testing.T) {
	c := config.ConfigFile{Path: "config_test.txt"}
	defer os.Remove("./config_test.txt")

	c.SetStore("http://foobar.test")

	got := c.GetConfig()
	want := config.Config{Store: "api", Path: "http://foobar.test"}
	if got != want {
		t.Errorf("got \"%s\", want \"%s\"", got, want)
	}
}

func Test_setting_file_store_config(t *testing.T) {
	c := config.ConfigFile{Path: "config_test.txt"}
	defer os.Remove("./config_test.txt")

	c.SetStore("foo.txt")

	got := c.GetConfig()
	want := config.Config{Store: "file", Path: "foo.txt"}
	if got != want {
		t.Errorf("got \"%s\", want \"%s\"", got, want)
	}
}

func Test_config_is_overwritten(t *testing.T) {
	c := config.ConfigFile{Path: "config_test.txt"}
	defer os.Remove("./config_test.txt")

	c.SetStore("http://foobar.test")
	c.SetStore("foo.txt")

	got := c.GetConfig()
	want := config.Config{Store: "file", Path: "foo.txt"}
	if got != want {
		t.Errorf("got \"%s\", want \"%s\"", got, want)
	}
}

func Test_store_config_uses_file_by_default(t *testing.T) {
	c := config.ConfigFile{Path: "config_test.txt"}
	defer os.Remove("./config_test.txt")

	got := c.GetConfig()
	want := config.Config{Store: "file", Path: "list.txt"}
	if got != want {
		t.Errorf("got \"%s\", want \"%s\"", got, want)
	}
}

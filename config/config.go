package config

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type ConfigFile struct {
	Path string
}

func (c ConfigFile) SetStore(store string, path string) {
	f, err := os.Create(c.Path)
	check(err)

	err = f.Truncate(0)
	check(err)

	_, err = f.Seek(0, 0)
	check(err)

	_, err = fmt.Fprintf(f, "%s\n%s", store, path)
	check(err)

	err = f.Close()
	check(err)
}

func (c ConfigFile) GetConfig() Config {
	_, err := os.Stat(c.Path)

	if err != nil {
		return Config{Store: "file", Path: "list.txt"}
	}

	b, err := os.ReadFile(c.Path)
	check(err)

	lines := strings.Split(string(b), "\n")
	return Config{Store: lines[0], Path: lines[1]}
}

type Config struct {
	Store string
	Path  string
}

package store

import (
	"bufio"
	"fmt"
	"os"
)

type Store interface {
	Add(item string)
	Next() string
}

type FileStore struct {
	Path string
}

func (s FileStore) Add(item string) {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}

	defer f.Close()

	f.WriteString(fmt.Sprintf("%s\n", item))
}

func (s FileStore) Next() string {
	file, err := os.Open(s.Path)

	if err != nil {
		return ""
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

type InMemoryStore []string

func (s *InMemoryStore) Add(item string) {
	*s = append(*s, item)
}

func (s *InMemoryStore) Next() string {
	if len(*s) == 0 {
		return ""
	}

	return (*s)[0]
}

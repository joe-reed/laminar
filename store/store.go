package store

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func check(err error) {
	if err != nil && err != io.EOF {
		panic(err)
	}
}

type Store interface {
	Add(item string)
	Next() string
	Pop() string
}

type FileStore struct {
	Path string
}

func (s FileStore) Add(item string) {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	check(err)

	defer f.Close()

	f.WriteString(fmt.Sprintf("%s\n", item))
}

func (s FileStore) Next() string {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	check(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		return scanner.Text()
	}

	return ""
}

func (s FileStore) Pop() string {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	check(err)

	fi, err := f.Stat()
	check(err)

	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	_, err = io.Copy(buf, f)
	check(err)

	line, err := buf.ReadBytes('\n')
	check(err)

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	nw, err := io.Copy(f, buf)
	check(err)

	err = f.Truncate(nw)
	check(err)

	err = f.Sync()
	check(err)

	_, err = f.Seek(0, io.SeekStart)
	check(err)

	if len(line) == 0 {
		return ""
	}

	return string(line[:len(line)-1])
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

func (s *InMemoryStore) Pop() string {
	item := s.Next()

	if len(*s) == 0 {
		return ""
	}

	*s = (*s)[1:]

	return item
}

type ApiStore struct {
	BaseURL string
	Client  *http.Client
}

func (s ApiStore) Add(item string) {
	_, err := s.Client.Post(s.BaseURL+"/add", "text/plain", strings.NewReader(item))

	if err != nil {
		panic(err)
	}
}

func (s ApiStore) Next() string {
	r, err := s.Client.Get(s.BaseURL + "/next")

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func (s ApiStore) Pop() string {
	r, err := s.Client.Get(s.BaseURL + "/pop")

	if err != nil {
		panic(err)
	}

	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	return string(b)
}

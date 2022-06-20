package store

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
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

func FromPath(path string) Store {
	if isUrl(path) {
		return NewApiStore(path)
	}

	return FileStore{Path: path}
}

func isUrl(s string) bool {
	_, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}

	u, err := url.Parse(s)
	return err == nil && u.Scheme != "" && u.Host != ""
}

type FileStore struct {
	Path string
}

func (s FileStore) Add(item string) {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", item))
	check(err)
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

	err = f.Close()
	check(err)

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
	BaseUrl string
	Client  *http.Client
}

func NewApiStore(baseUrl string) Store {
	return ApiStore{BaseUrl: baseUrl, Client: http.DefaultClient}
}

func (s ApiStore) Add(item string) {
	r, err := s.Client.Post(s.BaseUrl+"/add", "text/plain", strings.NewReader(item))
	check(err)

	if status := r.StatusCode; status != http.StatusCreated {
		panic(fmt.Sprintf("received status code %d", status))
	}
}

func (s ApiStore) Next() string {
	r, err := s.Client.Get(s.BaseUrl + "/next")
	check(err)

	if status := r.StatusCode; status != http.StatusOK {
		panic(fmt.Sprintf("received status code %d", status))
	}

	return responseToString(r)
}

func (s ApiStore) Pop() string {
	r, err := s.Client.Get(s.BaseUrl + "/pop")
	check(err)

	if status := r.StatusCode; status != http.StatusOK {
		panic(fmt.Sprintf("received status code %d", status))
	}

	return responseToString(r)
}

func responseToString(r *http.Response) string {
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)

	check(err)

	return string(b)
}

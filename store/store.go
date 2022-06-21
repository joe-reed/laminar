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

type Store interface {
	Add(item string) error
	Next() (string, error)
	Pop() (string, error)
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

func (s FileStore) Add(item string) error {
	f, err := os.OpenFile(s.Path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%s\n", item))
	if err != nil {
		return err
	}

	return nil
}

func (s FileStore) Next() (string, error) {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		return scanner.Text(), nil
	}

	return "", nil
}

func (s FileStore) Pop() (string, error) {
	f, err := os.OpenFile(s.Path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}

	fi, err := f.Stat()
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer(make([]byte, 0, fi.Size()))

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(buf, f)
	if err != nil {
		return "", err
	}

	line, err := buf.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	nw, err := io.Copy(f, buf)
	if err != nil {
		return "", err
	}

	err = f.Truncate(nw)
	if err != nil {
		return "", err
	}

	err = f.Sync()
	if err != nil {
		return "", err
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return "", err
	}

	if len(line) == 0 {
		return "", nil
	}

	err = f.Close()
	if err != nil {
		return "", err
	}

	return string(line[:len(line)-1]), nil
}

type InMemoryStore []string

func (s *InMemoryStore) Add(item string) error {
	*s = append(*s, item)

	return nil
}

func (s *InMemoryStore) Next() (string, error) {
	if len(*s) == 0 {
		return "", nil
	}

	return (*s)[0], nil
}

func (s *InMemoryStore) Pop() (string, error) {
	item, err := s.Next()
	if err != nil {
		return "", err
	}

	if len(*s) == 0 {
		return "", nil
	}

	*s = (*s)[1:]

	return item, nil
}

type ApiStore struct {
	BaseUrl string
	Client  *http.Client
}

func NewApiStore(baseUrl string) Store {
	return ApiStore{BaseUrl: baseUrl, Client: http.DefaultClient}
}

func (s ApiStore) Add(item string) error {
	r, err := s.Client.Post(s.BaseUrl+"/add", "text/plain", strings.NewReader(item))
	if err != nil {
		return err
	}

	if status := r.StatusCode; status != http.StatusCreated {
		return fmt.Errorf("received status code %d", status)
	}

	return nil
}

func (s ApiStore) Next() (string, error) {
	r, err := s.Client.Get(s.BaseUrl + "/next")
	if err != nil {
		return "", err
	}

	if status := r.StatusCode; status != http.StatusOK {
		return "", fmt.Errorf("received status code %d", status)
	}

	rs, err := responseToString(r)
	if err != nil {
		return "", err
	}

	return rs, nil
}

func (s ApiStore) Pop() (string, error) {
	r, err := s.Client.Get(s.BaseUrl + "/pop")
	if err != nil {
		return "", err
	}

	if status := r.StatusCode; status != http.StatusOK {
		return "", fmt.Errorf("received status code %d", status)
	}

	rs, err := responseToString(r)
	if err != nil {
		return "", err
	}

	return rs, nil
}

func responseToString(r *http.Response) (string, error) {
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

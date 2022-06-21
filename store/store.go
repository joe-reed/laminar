package store

import (
	"net/url"
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

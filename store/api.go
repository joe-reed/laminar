package store

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

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

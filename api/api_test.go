package api_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
)

func Test_add_returns_201(t *testing.T) {
	rr, _ := post(&store.InMemoryStore{}, "/add", "")

	assert.Equal(t, 201, rr.Code)
}

func Test_add_adds_an_item_to_the_store(t *testing.T) {
	store := &store.InMemoryStore{}

	post(store, "/add", "My next item")

	got, _ := store.Next()

	assert.Equal(t, "My next item", got)
}

func Test_next_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	assert.Equal(t, 200, rr.Code)
}

func Test_next_outputs_the_next_item(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	assert.Equal(t, "My next item", rr.Body.String())
}

func Test_next_outputs_an_empty_string_when_no_items_left(t *testing.T) {
	rr, _ := get(&store.InMemoryStore{}, "/next")

	assert.Equal(t, "", rr.Body.String())
}

func Test_pop_returns_200(t *testing.T) {
	rr, _ := get(&store.InMemoryStore{}, "/pop")

	assert.Equal(t, 200, rr.Code)
}

func Test_pop_removes_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")
	store.Add("My next item 2")

	get(store, "/pop")

	got, _ := store.Next()

	assert.Equal(t, "My next item 2", got)
}

func Test_pop_outputs_next_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")

	rr, _ := get(store, "/pop")

	assert.Equal(t, "My next item 1", rr.Body.String())
}

func Test_pop_outputs_empty_string_when_no_items_left(t *testing.T) {
	rr, _ := get(&store.InMemoryStore{}, "/pop")

	assert.Equal(t, "", rr.Body.String())
}

func Test_error_returned_as_500_response(t *testing.T) {
	store := errorStore{Error: "My error string"}

	tests := []struct {
		title string
		run   func() (*httptest.ResponseRecorder, error)
	}{
		{"pop", func() (*httptest.ResponseRecorder, error) { return get(store, "/pop") }},
		{"next", func() (*httptest.ResponseRecorder, error) { return get(store, "/next") }},
		{"add", func() (*httptest.ResponseRecorder, error) { return post(store, "/add", "") }},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			rr, _ := test.run()

			assert.Equal(t, 500, rr.Code)
			assert.Equal(t, "My error string\n", rr.Body.String())
		})
	}
}

func Test_invalid_methods_return_405_response(t *testing.T) {
	store := &store.InMemoryStore{}

	tests := []struct {
		title string
		run   func() (*httptest.ResponseRecorder, error)
	}{
		{"pop", func() (*httptest.ResponseRecorder, error) { return handleRequest(store, "/pop", http.MethodPost, "") }},
		{"next", func() (*httptest.ResponseRecorder, error) { return handleRequest(store, "/next", http.MethodPost, "") }},
		{"add", func() (*httptest.ResponseRecorder, error) { return handleRequest(store, "/add", http.MethodGet, "") }},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			rr, _ := test.run()

			assert.Equal(t, 405, rr.Code)
		})
	}
}

func get(s store.Store, url string) (*httptest.ResponseRecorder, error) {
	return handleRequest(s, url, http.MethodGet, "")
}

func post(s store.Store, url string, body string) (*httptest.ResponseRecorder, error) {
	return handleRequest(s, url, http.MethodPost, body)
}

func handleRequest(s store.Store, url string, method string, body string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	api.Handler(s).ServeHTTP(rr, req)

	return rr, nil
}

type errorStore struct {
	Error string
}

func (e errorStore) Add(item string) error {
	return errors.New(e.Error)
}

func (e errorStore) Next() (string, error) {
	return "", errors.New(e.Error)
}

func (e errorStore) Pop() (string, error) {
	return "", errors.New(e.Error)
}

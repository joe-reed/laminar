package api_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/store"
)

func Test_add_returns_201(t *testing.T) {
	store := &store.InMemoryStore{}

	rr, _ := post(store, "/add", "")

	if status := rr.Code; status != 201 {
		t.Errorf("got \"%d\" want \"%d\"", status, 201)
	}
}

func Test_add_adds_an_item_to_the_store(t *testing.T) {
	store := &store.InMemoryStore{}

	want := "My next item"

	post(store, "/add", want)

	got := store.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_next_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%d\" want \"%d\"", status, 200)
	}
}

func Test_next_outputs_the_next_item(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	want := "My next item"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_next_outputs_an_empty_string_when_no_items_left(t *testing.T) {
	store := &store.InMemoryStore{}

	rr, _ := get(store, "/next")

	want := ""
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_pop_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/pop")

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%d\" want \"%d\"", status, 200)
	}
}

func Test_pop_removes_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")
	store.Add("My next item 2")

	get(store, "/pop")

	want := "My next item 2"
	got := store.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_pop_outputs_next_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")

	rr, _ := get(store, "/pop")

	want := "My next item 1"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_pop_outputs_empty_string_when_no_items_left(t *testing.T) {
	store := &store.InMemoryStore{}

	rr, _ := get(store, "/pop")

	want := ""
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func get(s store.Store, url string) (*httptest.ResponseRecorder, error) {
	return handleRequest(s, url, http.MethodPost, "")
}

func post(s store.Store, url string, body string) (*httptest.ResponseRecorder, error) {
	return handleRequest(s, url, http.MethodPost, body)
}

func handleRequest(s store.Store, url string, method string, body string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(http.MethodGet, url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	api.Handler(s).ServeHTTP(rr, req)

	return rr, nil
}

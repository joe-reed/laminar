package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_getting_next_item(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_getting_next_item_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/next")

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("got %v want %v", status, http.StatusOK)
	}
}

func handleRequest(s store.Store, method string, url string) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()

	Handler(s).ServeHTTP(rr, req)

	return rr, nil
}

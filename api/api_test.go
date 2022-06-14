package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_getting_next_item_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_the_next_item(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_a_message_if_getting_next_item_when_all_items_complete(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_completing_item_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := handleRequest(store, "GET", "/done")

	want := "My next item\n"
	got := rr.Body.String()

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_completes_item_when_done(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")
	store.Add("My next item 2")

	handleRequest(store, "GET", "/done")

	want := "My next item 2"
	got := store.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_success_message_and_next_item_when_done(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")
	store.Add("My next item 2")

	rr, _ := handleRequest(store, "GET", "/done")

	want := "Item complete\nNext: My next item 2\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_a_message_when_completing_last_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")

	rr, _ := handleRequest(store, "GET", "/done")

	want := "Item complete\nAll items complete!\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
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

package api

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_adding_item_returns_201(t *testing.T) {
	store := &store.InMemoryStore{}

	rr, _ := post(store, "/add", "")

	if status := rr.Code; status != 201 {
		t.Errorf("got \"%d\" want \"%d\"", status, 201)
	}
}

func Test_it_adds_an_item_to_the_store(t *testing.T) {
	store := &store.InMemoryStore{}

	want := "My next item"

	post(store, "/add", want)

	got := store.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	store := &store.InMemoryStore{}

	rr, _ := post(store, "/add", "My next item")

	want := "Item added\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_getting_next_item_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%d\" want \"%d\"", status, 200)
	}
}

func Test_it_outputs_the_next_item(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_a_message_if_getting_next_item_when_all_items_complete(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/next")

	want := "My next item\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_completing_item_returns_200(t *testing.T) {
	store := &store.InMemoryStore{}
	store.Add("My next item")

	rr, _ := get(store, "/done")

	if status := rr.Code; status != 200 {
		t.Errorf("got \"%d\" want \"%d\"", status, 200)
	}
}

func Test_it_completes_item_when_done(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")
	store.Add("My next item 2")

	get(store, "/done")

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

	rr, _ := get(store, "/done")

	want := "Item complete\nNext: My next item 2\n"
	got := rr.Body.String()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_a_message_when_completing_last_item(t *testing.T) {
	store := &store.InMemoryStore{}

	store.Add("My next item 1")

	rr, _ := get(store, "/done")

	want := "Item complete\nAll items complete!\n"
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

	Handler(s).ServeHTTP(rr, req)

	return rr, nil
}

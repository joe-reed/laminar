package store_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/store"
)

func runSuite(t *testing.T, factory func() store.Store, teardown func()) {
	tests := []struct {
		title string
		run   func(t *testing.T, s store.Store)
	}{
		{"an added item is stored", testAddingItem},
		{"next returns empty string when no items left", testEmptyNext},
		{"pop removes the next item", testPopRemovesItem},
		{"pop returns the next item", testPopReturnsItem},
		{"pop returns empty string when no items left", testEmptyPop},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			test.run(t, factory())
			teardown()
		})
	}
}

func testAddingItem(t *testing.T, s store.Store) {
	expected := "My new item"
	s.Add(expected)
	actual := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyNext(t *testing.T, s store.Store) {
	actual := s.Next()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
	}
}

func testPopRemovesItem(t *testing.T, s store.Store) {
	expected := "Item 2"
	s.Add("Item 1")
	s.Add(expected)

	s.Pop()

	actual := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testPopReturnsItem(t *testing.T, s store.Store) {
	expected := "Item 1"
	s.Add(expected)
	s.Add("Item 2")

	actual := s.Pop()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyPop(t *testing.T, s store.Store) {
	actual := s.Pop()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
	}
}

func TestFileStore(t *testing.T) {
	runSuite(
		t,
		func() store.Store { return store.FileStore{"./list_test.txt"} },
		func() { os.Remove("./list_test.txt") },
	)
}

func TestInMemoryStore(t *testing.T) {
	runSuite(t, func() store.Store { return &store.InMemoryStore{} }, func() {})
}

func TestApiStore(t *testing.T) {
	s := store.InMemoryStore{}
	server := httptest.NewServer(api.Handler(&s))
	defer server.Close()

	runSuite(
		t,
		func() store.Store { return store.ApiStore{BaseUrl: server.URL, Client: server.Client()} },
		func() { s = store.InMemoryStore{} },
	)
}

func Test_api_store_panics_when_receiving_unexpected_status_code(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()

	s := store.ApiStore{BaseUrl: server.URL, Client: server.Client()}

	tests := []struct {
		title string
		run   func()
	}{
		{"pop", func() { s.Pop() }},
		{"next", func() { s.Next() }},
		{"add", func() { s.Add("foo") }},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			defer func() {
				got := recover()
				if got == nil {
					t.Errorf("The code did not panic")
				}

				want := "received status code 404"
				if got != want {
					t.Errorf("got \"%s\", want \"%s\"", got, want)
				}
			}()
			test.run()
		})
	}
}

package store

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func runSuite(t *testing.T, factory func() Store, teardown func()) {
	tests := []struct {
		title string
		run   func(t *testing.T, s Store)
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

func testAddingItem(t *testing.T, s Store) {
	expected := "My new item"
	s.Add(expected)
	actual := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyNext(t *testing.T, s Store) {
	actual := s.Next()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
	}
}

func testPopRemovesItem(t *testing.T, s Store) {
	expected := "Item 2"
	s.Add("Item 1")
	s.Add(expected)

	s.Pop()

	actual := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testPopReturnsItem(t *testing.T, s Store) {
	expected := "Item 1"
	s.Add(expected)
	s.Add("Item 2")

	actual := s.Pop()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyPop(t *testing.T, s Store) {
	actual := s.Pop()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
	}
}

func TestFileStore(t *testing.T) {
	runSuite(
		t,
		func() Store { return FileStore{"./test.txt"} },
		func() { os.Remove("./test.txt") },
	)
}

func TestInMemoryStore(t *testing.T) {
	runSuite(t, func() Store { return &InMemoryStore{} }, func() {})
}

func TestApiStore(t *testing.T) {
	store := InMemoryStore{}
	server := httptest.NewServer(handler(&store))
	defer server.Close()

	runSuite(
		t,
		func() Store { return ApiStore{BaseURL: server.URL, Client: server.Client()} },
		func() { store = InMemoryStore{} },
	)
}

func handler(s Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.Next())
	})

	mux.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.Pop())
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		s.Add(string(b))

		w.WriteHeader(http.StatusCreated)
	})

	return mux
}

package store_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joe-reed/laminar/api"
	"github.com/joe-reed/laminar/store"
)

func Test_api_store(t *testing.T) {
	s := store.InMemoryStore{}
	server := httptest.NewServer(api.Handler(&s))
	defer server.Close()

	runSuite(
		t,
		func() store.Store { return store.ApiStore{BaseUrl: server.URL, Client: server.Client()} },
		func() { s = store.InMemoryStore{} },
	)
}

func Test_api_store_returns_error_when_receiving_unexpected_status_code(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}))
	defer server.Close()

	s := store.ApiStore{BaseUrl: server.URL, Client: server.Client()}

	tests := []struct {
		title string
		run   func() error
	}{
		{"pop", func() error { _, err := s.Pop(); return err }},
		{"next", func() error { _, err := s.Next(); return err }},
		{"add", func() error { return s.Add("foo") }},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got := test.run().Error()
			want := "received status code 404"

			if got != want {
				t.Errorf("got \"%s\", want \"%s\"", got, want)
			}
		})
	}
}

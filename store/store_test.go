package store_test

import (
	"testing"

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
	want := "My new item"
	s.Add(want)
	got, _ := s.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func testEmptyNext(t *testing.T, s store.Store) {
	got, _ := s.Next()
	want := ""

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func testPopRemovesItem(t *testing.T, s store.Store) {
	want := "Item 2"
	s.Add("Item 1")
	s.Add(want)

	s.Pop()

	got, _ := s.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func testPopReturnsItem(t *testing.T, s store.Store) {
	want := "Item 1"
	s.Add(want)
	s.Add("Item 2")

	got, _ := s.Pop()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func testEmptyPop(t *testing.T, s store.Store) {
	got, _ := s.Pop()
	want := ""

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_getting_store_from_path(t *testing.T) {
	tests := []struct {
		title string
		path  string
		store store.Store
	}{
		{"file", "foo.txt", store.FileStore{Path: "foo.txt"}},
		{"api", "http://foo.test", store.NewApiStore("http://foo.test")},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			got := store.FromPath(test.path)
			want := test.store

			if got != want {
				t.Errorf("got \"%s\" want \"%s\"", got, want)
			}
		})
	}
}

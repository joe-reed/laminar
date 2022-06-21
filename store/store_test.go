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
	expected := "My new item"
	s.Add(expected)
	actual, _ := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyNext(t *testing.T, s store.Store) {
	actual, _ := s.Next()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
	}
}

func testPopRemovesItem(t *testing.T, s store.Store) {
	expected := "Item 2"
	s.Add("Item 1")
	s.Add(expected)

	s.Pop()

	actual, _ := s.Next()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testPopReturnsItem(t *testing.T, s store.Store) {
	expected := "Item 1"
	s.Add(expected)
	s.Add("Item 2")

	actual, _ := s.Pop()

	if expected != actual {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyPop(t *testing.T, s store.Store) {
	actual, _ := s.Pop()

	if actual != "" {
		t.Errorf("Expected %s, got %s", "", actual)
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
				t.Errorf("got \"%s\", want \"%s\"", got, want)
			}
		})
	}
}

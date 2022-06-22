package store_test

import (
	"testing"

	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
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
	s.Add("My new item")

	got, _ := s.Next()

	assert.Equal(t, "My new item", got)
}

func testEmptyNext(t *testing.T, s store.Store) {
	got, _ := s.Next()
	assert.Equal(t, "", got)
}

func testPopRemovesItem(t *testing.T, s store.Store) {
	s.Add("Item 1")
	s.Add("Item 2")

	s.Pop()

	got, _ := s.Next()

	assert.Equal(t, "Item 2", got)
}

func testPopReturnsItem(t *testing.T, s store.Store) {
	s.Add("Item 1")
	s.Add("Item 2")

	got, _ := s.Pop()

	assert.Equal(t, "Item 1", got)
}

func testEmptyPop(t *testing.T, s store.Store) {
	got, _ := s.Pop()
	assert.Equal(t, "", got)
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
			assert.Equal(t, test.store, store.FromPath(test.path))
		})
	}
}

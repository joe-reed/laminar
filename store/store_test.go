package store

import (
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

	if expected != string(actual) {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}

func testEmptyNext(t *testing.T, s Store) {
	actual := s.Next()

	if "" != string(actual) {
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
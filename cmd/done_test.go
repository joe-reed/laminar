package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_it_completes_item_when_done(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")
	s.Add("My next item 2")

	runDone(t, s)

	want := "My next item 2"
	got := s.Next()

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func Test_it_outputs_success_message_and_next_item_when_done(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")
	s.Add("My next item 2")

	want := "Item complete\nNext: My next item 2\n"
	got := runDone(t, s)

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func Test_it_outputs_a_message_when_completing_last_item(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")

	want := "Item complete\nAll items complete!\n"
	got := runDone(t, s)

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func runDone(t *testing.T, s store.Store) string {
	return runCommand(t, s, []string{"done"})
}
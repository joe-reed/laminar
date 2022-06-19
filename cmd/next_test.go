package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_it_outputs_the_next_item(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item")

	want := "My next item\n"
	got := runNext(t, s)

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func Test_it_outputs_a_message_if_getting_next_item_when_all_items_complete(t *testing.T) {
	s := &store.InMemoryStore{}

	want := "All items complete!\n"
	got := runNext(t, s)

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func runNext(t *testing.T, s store.Store) string {
	return runCommand(t, s, []string{"next"})
}

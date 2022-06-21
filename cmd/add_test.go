package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
)

func Test_it_adds_an_item_to_the_store(t *testing.T) {
	s := &store.InMemoryStore{}

	want := "Test item"
	runAdd(t, s, want)

	got, _ := s.Next()

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	want := "Item added\n"
	got := runAdd(t, &store.InMemoryStore{}, "Test item")

	if got != want {
		t.Errorf("got \"%s\" want \"%s\"", got, want)
	}
}

func runAdd(t *testing.T, s store.Store, item string) string {
	return runCommand(t, s, config.New(), []string{"add", item})
}

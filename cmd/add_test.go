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

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	want := "Item added\n"
	got := runAdd(t, &store.InMemoryStore{}, "Test item")

	if want != got {
		t.Errorf("want \"%s\", got \"%s\"", want, got)
	}
}

func runAdd(t *testing.T, s store.Store, item string) string {
	return runCommand(t, s, config.New(), []string{"add", item})
}

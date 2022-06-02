package cli

import (
	"bytes"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_it_adds_an_item_to_the_store(t *testing.T) {
	store := store.InMemoryStore{}

	expected := "Refactor foo"

	Add(expected, &store, &bytes.Buffer{})

	actual := store.Next()

	if expected != actual {
		t.Errorf("Expected success message \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	store := store.InMemoryStore{}
	var output bytes.Buffer

	Add("Refactor foo", &store, &output)

	expected := "Item added\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected success message \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_the_next_item(t *testing.T) {
	store := store.InMemoryStore{}
	var output bytes.Buffer

	store.Add("My next item")

	Next(&store, &output)

	expected := "My next item\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected success message \"%s\", got \"%s\"", expected, actual)
	}
}

package cli

import (
	"bytes"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_it_adds_an_item_to_the_store(t *testing.T) {
	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &bytes.Buffer{},
	}

	expected := "Refactor foo"

	c.Add(expected)

	actual := c.Store.Next()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	output := bytes.Buffer{}

	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &output,
	}

	c.Add("Refactor foo")

	expected := "Item added\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_the_next_item(t *testing.T) {
	output := bytes.Buffer{}

	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &output,
	}

	c.Store.Add("My next item")

	c.Next()

	expected := "My next item\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_a_message_if_getting_next_item_when_all_items_complete(t *testing.T) {
	output := bytes.Buffer{}

	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &output,
	}

	c.Next()

	expected := "All items complete!\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_completes_item_when_done(t *testing.T) {
	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &bytes.Buffer{},
	}

	c.Store.Add("My next item 1")
	c.Store.Add("My next item 2")

	c.Done()

	expected := "My next item 2"
	actual := c.Store.Next()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_success_message_and_next_item_when_done(t *testing.T) {
	output := bytes.Buffer{}

	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &output,
	}

	c.Store.Add("My next item 1")
	c.Store.Add("My next item 2")

	c.Done()

	expected := "Item complete\nNext: My next item 2\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

func Test_it_outputs_a_message_when_completing_last_item(t *testing.T) {
	output := bytes.Buffer{}

	c := Cli{
		Store:  &store.InMemoryStore{},
		Output: &output,
	}

	c.Store.Add("My next item 1")

	c.Done()

	expected := "Item complete\nAll items complete!\n"
	actual := output.String()

	if expected != actual {
		t.Errorf("Expected \"%s\", got \"%s\"", expected, actual)
	}
}

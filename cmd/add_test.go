package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
)

func Test_it_adds_an_item_to_the_store(t *testing.T) {
	s := &store.InMemoryStore{}

	runAdd(t, s, "Test item")

	got, _ := s.Next()

	assert.Equal(t, "Test item", got)
}

func Test_it_outputs_success_message_when_adding_item(t *testing.T) {
	assert.Equal(
		t,
		"Item added\n",
		runAdd(t, &store.InMemoryStore{}, "Test item"),
	)
}

func runAdd(t *testing.T, s store.Store, item string) string {
	return runCommand(t, s, config.New(), []string{"add", item})
}

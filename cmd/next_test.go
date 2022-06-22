package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
)

func Test_it_outputs_the_next_item(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item")

	assert.Equal(t, "My next item\n", runNext(t, s))
}

func Test_it_outputs_a_message_if_getting_next_item_when_all_items_complete(t *testing.T) {
	assert.Equal(
		t,
		"All items complete!\n",
		runNext(t, &store.InMemoryStore{}),
	)
}

func runNext(t *testing.T, s store.Store) string {
	return runCommand(t, s, config.New(), []string{"next"})
}

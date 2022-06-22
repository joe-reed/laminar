package cmd_test

import (
	"testing"

	"github.com/joe-reed/laminar/config"
	"github.com/joe-reed/laminar/store"
	"github.com/stretchr/testify/assert"
)

func Test_it_completes_item_when_done(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")
	s.Add("My next item 2")

	runDone(t, s)

	got, _ := s.Next()

	assert.Equal(t, "My next item 2", got)
}

func Test_it_outputs_success_message_and_next_item_when_done(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")
	s.Add("My next item 2")

	assert.Equal(t, "Completed: My next item 1\nNext: My next item 2\n", runDone(t, s))
}

func Test_it_outputs_a_message_when_completing_last_item(t *testing.T) {
	s := &store.InMemoryStore{}

	s.Add("My next item 1")

	assert.Equal(t, "Completed: My next item 1\nAll items complete!\n", runDone(t, s))
}

func Test_it_does_not_output_success_message_when_all_items_complete(t *testing.T) {
	assert.Equal(t, "All items complete!\n", runDone(t, &store.InMemoryStore{}))
}

func runDone(t *testing.T, s store.Store) string {
	return runCommand(t, s, config.New(), []string{"done"})
}

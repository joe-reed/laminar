package store_test

import (
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_in_memory_store(t *testing.T) {
	runSuite(t, func() store.Store { return &store.InMemoryStore{} }, func() {})
}

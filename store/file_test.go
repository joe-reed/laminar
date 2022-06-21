package store_test

import (
	"os"
	"testing"

	"github.com/joe-reed/laminar/store"
)

func Test_file_store(t *testing.T) {
	runSuite(
		t,
		func() store.Store { return store.FileStore{"./list_test.txt"} },
		func() { os.Remove("./list_test.txt") },
	)
}

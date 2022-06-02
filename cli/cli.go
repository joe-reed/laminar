package cli

import (
	"fmt"
	"io"

	"github.com/joe-reed/laminar/store"
)

func Add(text string, store store.Store, output io.Writer) {
	store.Add(text)

	fmt.Fprintln(output, "Item added")
}

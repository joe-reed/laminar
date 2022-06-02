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

func Next(store store.Store, output io.Writer) {
	fmt.Fprintln(output, store.Next())
}

func Done(store store.Store, output io.Writer) {
	store.Pop()
	fmt.Fprintln(output, "Item complete")
	fmt.Fprintln(output, fmt.Sprintf("Next: %s", store.Next()))
}

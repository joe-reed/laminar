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
	if store.Next() == "" {
		fmt.Fprintln(output, "All items complete!")
		return
	}

	fmt.Fprintln(output, store.Next())
}

func Done(store store.Store, output io.Writer) {
	store.Pop()
	fmt.Fprintln(output, "Item complete")

	if store.Next() == "" {
		fmt.Fprintln(output, "All items complete!")
		return
	}

	fmt.Fprintf(output, "Next: %s\n", store.Next())
}

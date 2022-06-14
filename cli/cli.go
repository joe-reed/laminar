package cli

import (
	"fmt"
	"io"

	"github.com/joe-reed/laminar/store"
)

type Cli struct {
	Store  store.Store
	Output io.Writer
}

func (c Cli) Add(text string) {
	c.Store.Add(text)

	fmt.Fprintln(c.Output, "Item added")
}

func (c Cli) Next() {
	if c.Store.Next() == "" {
		fmt.Fprintln(c.Output, "All items complete!")
		return
	}

	fmt.Fprintln(c.Output, c.Store.Next())
}

func (c Cli) Done() {
	c.Store.Pop()
	fmt.Fprintln(c.Output, "Item complete")

	if c.Store.Next() == "" {
		fmt.Fprintln(c.Output, "All items complete!")
		return
	}

	fmt.Fprintf(c.Output, "Next: %s\n", c.Store.Next())
}

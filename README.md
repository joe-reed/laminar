# Laminar

CLI todo list for focus and flow.

## Motivation
Laminar is a FIFO task list that helps you stay focused on the next task at hand.
Use cases:
- Focus on the next test in a TDD loop, adding new cases that you think of to the list for later
- Avoid distraction as a driver during pair programming by grabbing the next thing on the todo list, while the navigator adds to the list remotely
- Any other list you can think of that would benefit from working through things one at a time!

## Setup
- Install [Go](https://go.dev/)
- `make install`

## Usage
Add a new item to your list using `add`
```
laminar add "Your new item"
```

See what's next on the list using `next`
```
laminar next

// Your new item
```

Complete an item using `done`
```
laminar done

// Item complete
// Next: take out the bins
```

Use `serve` to make your local Laminar accessible via an API:
```
laminar serve
```

Configure your Laminar to use an API from `serve` as its back end:
```
laminar configure http://url-from-laminar-serve.test
```

Or point it at a local file:
```
laminar configure /path/to/file.txt
```

For full usage instructions, run
```
laminar --help
```

## Tests
`make test`

## Build locally
`make build`

## Dependencies
- [Cobra](https://github.com/spf13/cobra) for the CLI commands
- [Viper](https://github.com/spf13/viper) for managing configuration
- [go-localtunnel](https://github.com/localtunnel/go-localtunnel) for exposing the local API
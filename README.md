# Laminar

CLI todo list for focus and flow.

## Setup
- Install [Go](https://go.dev/)
- `make build`

## Usage
Add a new item to your list using `add`
```
./bin/laminar add "Your new item"
```

See what's next on the list using `next`
```
./bin/laminar next

// Your new item
```

Complete an item using `done`
```
./bin/laminar done

// Item complete
// Next: take out the bins
```

Use `serve` to make your local Laminar accessible via an API:
```
./bin/laminar serve
```

Configure your Laminar to use an API from `serve` as its back end:
```
./bin/laminar configure http://url-from-laminar-serve.test
```

Or point it at a local file:
```
./bin/laminar configure /path/to/file.txt
```

## Tests
`make test`
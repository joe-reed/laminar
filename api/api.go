package api

import (
	"io"
	"log"
	"net"
	"net/http"

	"github.com/joe-reed/laminar/cli"
	"github.com/joe-reed/laminar/store"
	"github.com/localtunnel/go-localtunnel"
)

func Serve(s store.Store) {

	server := &http.Server{Handler: Handler(s)}

	err := server.Serve(Listener())

	if err != nil {
		panic(err)
	}
}

func Handler(s store.Store) http.Handler {
	c := func(w io.Writer) cli.Cli {
		return cli.Cli{Store: s, Output: w}
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		c(w).Next()
	})

	mux.HandleFunc("/done", func(w http.ResponseWriter, r *http.Request) {
		c(w).Done()
	})

	return mux
}

func Listener() net.Listener {
	listener, err := localtunnel.Listen(localtunnel.Options{Log: log.Default()})

	if err != nil {
		panic(err)
	}

	return listener
}

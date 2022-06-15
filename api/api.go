package api

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

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
	mux := http.NewServeMux()

	mux.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.Next())
	})

	mux.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, s.Pop())
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)

		if err != nil {
			panic(err)
		}

		s.Add(string(b))

		w.WriteHeader(http.StatusCreated)
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

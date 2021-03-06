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

func Serve(s store.Store) error {
	server := &http.Server{Handler: Handler(s)}

	listener, err := Listener()
	if err != nil {
		return err
	}

	return server.Serve(listener)
}

func Handler(s store.Store) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			methodNotAllowed(w)
			return
		}

		next, err := s.Next()
		if err != nil {
			serverError(w, err)
			return
		}

		fmt.Fprint(w, next)
	})

	mux.HandleFunc("/pop", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			methodNotAllowed(w)
			return
		}

		next, err := s.Pop()
		if err != nil {
			serverError(w, err)
			return
		}

		fmt.Fprint(w, next)
	})

	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			methodNotAllowed(w)
			return
		}

		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)
		if err != nil {
			serverError(w, err)
			return
		}

		err = s.Add(string(b))
		if err != nil {
			serverError(w, err)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	return mux
}

func Listener() (net.Listener, error) {
	return localtunnel.Listen(localtunnel.Options{Log: log.Default()})
}

func serverError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func methodNotAllowed(w http.ResponseWriter) {
	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

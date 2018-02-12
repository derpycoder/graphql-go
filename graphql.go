package main

import (
	"net/http"

	"github.com/graphql-go/handler"
	"github.com/graphql-go/relay/examples/starwars"
)

func init() {
	h := handler.New(&handler.Config{
		Schema: &starwars.Schema,
		Pretty: true,
	})

	http.Handle("/", withCORS(h.ServeHTTP))
}

// Simple wrapper to Allow CORS.
func withCORS(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		fn(w, r)
	}
}

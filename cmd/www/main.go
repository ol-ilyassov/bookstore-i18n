package main

import (
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

func main() {
	mux := pat.New()
	mux.Get("/:locale", http.HandlerFunc(handleHome))

	log.Print("starting server on :4018...")
	err := http.ListenAndServe(":4018", mux)
	log.Fatal(err)
}

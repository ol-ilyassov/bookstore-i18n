package main

import (
	"fmt"
	"net/http"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Extract the locale from the URL path.
	locale := r.URL.Query().Get(":locale")

	switch locale {
	case "en-gb", "de-de", "fr-ch":
		fmt.Fprintf(w, "The locale is %s\n", locale)
	default:
		http.NotFound(w, r)
	}
}

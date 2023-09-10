package main

import (
	"net/http"

	// Import the internal/translations package, so that its init() function is called.
	_ "bookstore.example.com/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Extract the locale from the URL path.
	locale := r.URL.Query().Get(":locale")

	// Declare variable to hold the target language tag.
	var lang language.Tag

	switch locale {
	case "en-gb":
		lang = language.MustParse("en-GB")
	case "de-de":
		lang = language.MustParse("de-DE")
	case "fr-ch":
		lang = language.MustParse("fr-CH")
	default:
		http.NotFound(w, r)
		return
	}

	var totalBookCount = 1_252_794

	// Initialize a message.Printer which uses the target language.
	p := message.NewPrinter(lang)
	// Print the welcome message translated into the target language.
	p.Fprintf(w, "Welcome!\n")

	// Example of translated message with enterpolated integer value.
	p.Fprintf(w, "%d books available\n", totalBookCount)
}

package localizer

import (
	// important import for catalog messages.
	_ "bookstore.example.com/internal/translations"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// Define a Localizer type which stores the relevant locale ID (as used
// in our URLs) and a (deliberately unexported) message.Printer instance
// for the locale.
type Localizer struct {
	ID      string
	printer *message.Printer
}

// Initialize a slice which holds the initialized Localizer types for
// each of our supported locales.
var locales = []Localizer{
	{
		// Germany
		ID:      "de-de",
		printer: message.NewPrinter(language.MustParse("de-DE")),
	},
	{
		// Switzerland (French speaking)
		ID:      "fr-ch",
		printer: message.NewPrinter(language.MustParse("fr-CH")),
	},
	{
		// United Kingdom
		ID:      "en-gb",
		printer: message.NewPrinter(language.MustParse("en-GB")),
	},
}

// The Get() function accepts a locale ID and returns the corresponding
// Localizer for that locale. If the locale ID is not supported then
// this returns `false` as the second return value.
func Get(id string) (Localizer, bool) {
	for _, locale := range locales {
		if id == locale.ID {

			return locale, true
		}
	}

	return Localizer{}, false
}

// We also add a Translate() method to the Localizer type. This acts
// as a wrapper around the unexported message.Printer's Sprintf()
// function and returns the appropriate translation for the given
// message and arguments.
func (l Localizer) Translate(key message.Reference, args ...interface{}) string {
	return l.printer.Sprintf(key, args...)
}

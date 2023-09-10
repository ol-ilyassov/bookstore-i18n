
.SILENT:

.PHONY: run
run:
	go run ./cmd/www/

.PHONY: generate
generate:
	go generate ./internal/translations/translations.go

.PHONY: copy
copy:
	cp internal/translations/locales/de-DE/out.gotext.json internal/translations/locales/de-DE/messages.gotext.json
	cp internal/translations/locales/fr-CH/out.gotext.json internal/translations/locales/fr-CH/messages.gotext.json
	cp internal/translations/locales/en-GB/out.gotext.json internal/translations/locales/en-GB/messages.gotext.json
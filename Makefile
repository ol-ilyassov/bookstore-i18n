
.SILENT:

## make run: runs project as HTTP server.
.PHONY: run
run:
	go run ./cmd/www/

## make generate: generates translation message files for specified languges.
.PHONY: generate
generate:
	go generate ./internal/translations/translations.go

## make copy: copies translation message files to required format.
.PHONY: copy
copy:
	cp internal/translations/locales/de-DE/out.gotext.json internal/translations/locales/de-DE/messages.gotext.json
	cp internal/translations/locales/fr-CH/out.gotext.json internal/translations/locales/fr-CH/messages.gotext.json
	cp internal/translations/locales/en-GB/out.gotext.json internal/translations/locales/en-GB/messages.gotext.json
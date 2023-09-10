
.SILENT:

.PHONY: run
run:
	go run ./cmd/www/

.PHONY: generate
generate:
	go generate ./internal/translations/translations.go
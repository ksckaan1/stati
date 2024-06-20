generate:
	@templ generate

run: build
	@go run ./example/

dev:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="go run ./example/"
generate-templ:
	@templ generate

generate-css:
	@npx tailwindcss -i ./templates/style/input.css -o ./assets/style.css 

build: generate-templ generate-css

run: generate-css
	@go run ./example/

dev:
	@templ generate --watch --proxy="http://localhost:3000" --cmd="make run" --open-browser=false
dev:
	npx tailwindcss -i ./assets/css/tailwind.css -o ./assets/dist/styles.css
	TEMPL_EXPERIMENT=rawgo templ generate
	go run cmd/main.go

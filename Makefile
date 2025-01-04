run : build
	@./bin/main

build:
	@air

css:
	@./tailwindcss -i cmd/web/assets/css/tailwind.css -o cmd/web/assets/css/index.css




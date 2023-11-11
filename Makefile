build:
	mkdir -p bin
	@rm -rf main

	@CGO_ENABLED=1
	@go build main.go -o ./bin/main

run:build
	./bin/main

watch:
	find ./ -type f -iname "*.go" -o -iname "*.html" | entr make run
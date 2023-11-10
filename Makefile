build:
	@rm -rf main
	@CGO_ENABLED=1 
	@go build main.go

run:build
	./main

watch:
	find ./ -type f -iname "*.go" -o -iname "*.html" | entr make run
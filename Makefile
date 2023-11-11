build:
	@rm -rf ./bin
	mkdir -p bin/db
	cp -r static ./bin/
	cp -r templates ./bin/
	@rm -rf main

	@CGO_ENABLED=1 go build -o ./bin/gohtmx ./cmd/gohtmx/gohtmx.go

run:build
	./bin/gohtmx

watch:
	find ./ -type f -iname "*.go" -o -iname "*.html" | entr make run
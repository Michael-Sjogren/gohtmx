MAIN_FILE=./cmd/main/main.go
PROGRAM_NAME=gohtmx.out
BUILD_LOC=./bin
PROGRAM_PATH=$(BUILD_LOC)/$(PROGRAM_NAME)

build:
	@rm -rf $(BUILD_LOC)
	mkdir -p $(BUILD_LOC)
	@echo "Copying files...\n"
	cp -r static $(BUILD_LOC)
	cp -r templates $(BUILD_LOC)
	@echo "Building...\n"
	CGO_ENABLED=1 go build -o $(PROGRAM_PATH) $(MAIN_FILE)

run:build
	$(PROGRAM_PATH)

watch:
	find ./ -type f -iname "*.go" -o -iname "*.html" | entr make run
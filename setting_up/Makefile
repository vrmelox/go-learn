.PHONY : build clean run help test

SRC = read_number_ligne.go

BINARY_NAME = setting_up

build: 
	go build -o $(BINARY_NAME) $(SRC)

clean: 
	rm -rf $(BINARY_NAME)
	go clean

run: build
	./$(BINARY_NAME)

help:
	@echo "Available targets:"
	@echo "  build   - Compile the Go application"
	@echo "  test    - Run all Go tests"
	@echo "  clean   - Remove compiled binaries and other artifacts"
	@echo "  run     - Build and run the application"
	@echo "  help    - Display this help message"
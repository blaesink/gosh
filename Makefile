all: prep exe

prep:
	rm -rf gosh
	clear
	@echo "Running tests..."
	@echo
	@echo
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

exe:
	go build .

.PHONY: all

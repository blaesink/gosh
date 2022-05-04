all: prep exe


prep:
	rm -rf gosh
	clear
	@echo "Running tests..."
	@echo
	@echo
	go test ./... -v

exe:
	go build .

.PHONY: all

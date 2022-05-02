all: prep exe


prep:
	rm -rf gosh
	go test ./...

exe:
	go build .

.PHONY: all

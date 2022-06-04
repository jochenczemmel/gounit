
COVERFILE=".coverage.txt"

test:
	go test ./...

testv:
	go test -v ./...


cover:
	go test -cover -coverprofile $(COVERFILE) ./...
	go tool cover -html=$(COVERFILE)


.PHONY: test testv cover



 
.PHONEY: all install update build run

# Setup tools
setup:
	go get github.com/tools/godep
	go get golang.org/x/tools/cmd/cover

# Install and restore dependencies
restore:
	godep restore

# Update dependencies
update:
	go get -u ./...
	godep save

# Build script
build:
	go build -o ./server ./api/*.go

install: build
	go install ./...

run: build
	./server

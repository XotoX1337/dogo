icon:
	rsrc -ico dogo.ico

build:
	go mod tidy
	go build -ldflags "-s -w" -o bin/ .

run:
	go run .

clean:
	go clean

compile:
	go mod tidy
	GOOS=linux GOARCH=arm go build -ldflags "-s -w" -o ./dogo-linux-arm .
	GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./dogo-linux-amd64 .
	GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o ./dogo-windows-amd64 .
	GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o ./dogo-macos-amd64 .

all: clean compile
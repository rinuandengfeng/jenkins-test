
mod:
	go mod tidy
build:
	CGO_ENABLE=1 GOARCH=amd64 GOOS=linux go build -a -o .

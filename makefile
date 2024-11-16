
.PHONY:all
all: mod build

.PHONY:mod
mod:
	go mod tidy

.PHONY:build
build:
	CGO_ENABLE=1 GOARCH=amd64 GOOS=linux go build -a . -o app

clean:
	rm -f app


dev:
	air -c config/.air.toml

test:
	go test -cover

build:
	go build -o bin/boilerplate-go
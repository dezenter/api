dev:
	air

run:
	go run main.go

build:
	go build main.go

clean-cache:
	go clean -modcache

update:
	go get -u
dev:
	air

run:
	go run main.go

build:
	go build main.go

ccache:
	go clean -modcache

update:
	go get -u
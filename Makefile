all: clean build

clean:
	rm -f ./examples/basic/basic
	rm -f ./examples/cli-service/cli-service
	rm -f ./examples/high-priority/high-priority
	rm -f ./examples/locked-priority/locked-priority
	rm -f ./examples/lazy-logger/lazy-logger

build:
	cd ./examples/basic && go build
	cd ./examples/cli-service && go build
	cd ./examples/high-priority && go build
	cd ./examples/locked-priority && go build
	cd ./examples/lazy-logger && go build

build-race:
	cd ./examples/basic && go build --race
	cd ./examples/cli-service && go build --race
	cd ./examples/high-priority && go build --race
	cd ./examples/locked-priority && go build --race
	cd ./examples/lazy-logger && go build --race

test: #run all tests
	go test ./...

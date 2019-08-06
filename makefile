build: main
	
test-run: clean build
	docker run golang-test

clean:
	docker run --rm -v ${PWD}:/work -w /work alpine rm -rf bin

package = /go/src/github.com/tahomatx/hello-golang

init:
	docker build -f ./Dockerfile.dev -t golang-dev .

dev:
	docker run --rm -it -v $(PWD):$(package) -w $(package) golang-dev

main:
	docker run --rm -v $(PWD):$(package) -w $(package) golang-dev go build -o bin/main
	docker build -t golang-test .

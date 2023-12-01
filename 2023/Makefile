all: run

build:
	go build -o aoc2023

format:
	go fmt ./...

run: format build
	./aoc2023 $(ARGS)
	
init:
	cp -r tmpl day$(ARGS)
	sed -i '' -e 's/X/$(ARGS)/g' day$(ARGS)/solution.go

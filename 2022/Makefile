all: run

build:
	go build -o aoc2022

format:
	go fmt ./...

run: format build
	./aoc2022 $(ARGS)
	
init:
	cp -r tmpl day$(ARGS)
	sed -i 's/X/$(ARGS)/g' day$(ARGS)/solution.go

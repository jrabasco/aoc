all: run

build:
	go build -o aoc2023

format:
	go fmt ./...

LIST = 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
define make-day-target
day$1: format build
	./aoc2023 -day $1
endef

define make-day-profile-target
profile-day$1: format build
	./aoc2023 -day $1 -cpuprofile aoc2023.prof
	go tool pprof aoc2023 aoc2023.prof
endef

$(foreach element,$(LIST),$(eval $(call make-day-target,$(element))))
$(foreach element,$(LIST),$(eval $(call make-day-profile-target,$(element))))

run: format build
	./aoc2023

init:
	cp -r tmpl day$(ARGS)
	sed -i '' -e 's/X/$(ARGS)/g' day$(ARGS)/solution.go

test: format build
	./aoc2023 -test

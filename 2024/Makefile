all: run

build:
	go build -o aoc2024

format:
	go fmt ./...

LIST = 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
define make-day-target
day$1: format build
	./aoc2024 -day $1
endef

define make-day-profile-target
profile-day$1: format build
	./aoc2024 -day $1 -cpuprofile aoc2024.prof
	go tool pprof aoc2024 aoc2024.prof
endef

define make-day-time-target
time-day$1: format build
	./aoc2024 -day $1 -time
endef

define make-init-day
init-day$1:
	cp -r tmpl day$1
	sed -i '' -e 's/X/$1/g' day$1/solution.go
endef

$(foreach element,$(LIST),$(eval $(call make-day-target,$(element))))
$(foreach element,$(LIST),$(eval $(call make-day-profile-target,$(element))))
$(foreach element,$(LIST),$(eval $(call make-day-time-target,$(element))))
$(foreach element,$(LIST),$(eval $(call make-init-day,$(element))))

run: format build
	./aoc2024

test: format build
	./aoc2024 -test

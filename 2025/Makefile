all: run

build:
	CGO_CFLAGS="-I/opt/homebrew/include" CGO_LDFLAGS="-L/opt/homebrew/lib" go build -o aoc2025

format:
	go fmt ./...

LIST = 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
define make-day-target
day$1: format build
	./aoc2025 -day $1
endef

define make-day-profile-target
profile-day$1: format build
	./aoc2025 -day $1 -cpuprofile aoc2025.prof
	go tool pprof aoc2025 aoc2025.prof
endef

define make-day-time-target
time-day$1: format build
	./aoc2025 -day $1 -time
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
	./aoc2025

test: format build
	./aoc2025 -test

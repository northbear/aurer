
# SOURCES=(aurer.go args.go)

.PHONY: all test build clean

all: test build

test:
	go test

build: aurer.go config.go
	[[ -d build ]] || mkdir -p build/bin; \
	go build -o build/bin/aurer $^

install: aurer.go config.go
	[[ -d build ]] || mkdir -p build/bin; \
	go install $^

clean:
	[[ -d build ]] && rm -rf build


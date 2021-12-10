
# SOURCES=(aurer.go args.go)

.PHONY: all test build clean

all: build test

test:
	go test

build: aurer.go config.go
	[[ -d build ]] || mkdir -p build/bin
	go build -o build/bin/aurer $^

clean:
	[[ -d build ]] && rm -rf build


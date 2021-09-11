
# SOURCES=(aurer.go args.go)

.PHONY: all test build clean

all:

test:

build: aurer.go
	[[ -d build ]] || mkdir -p build/bin
	go build -o build/bin/aurer $^

clean:
	[[ -d build ]] && rm -rf build

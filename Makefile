run: build
	bin/random

build: clean
	go build -o bin/random

clean:
	rm -f bin/random

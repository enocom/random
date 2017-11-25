run: build
	bin/random

build: clean
	go build -o bin/random

clean:
	rm -f bin/random

docker: clean docker-bin
	docker build -t random-app

docker-bin:
	CGO_ENABLED=0 \
	GOOS=linux \
	GOARCH=amd64 \
	    go build \
	    -a \
	    -installsuffix nocgo \
	    -o ./bin/random \
	    github.com/enocom/random

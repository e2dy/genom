.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/genom-project/genom/cmd/swarm
	go build -o bin/geth     github.com/genom-project/genom/cmd/geth
	go build -o bin/bootnode github.com/genom-project/genom/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/genom-project/genom/swarm/...

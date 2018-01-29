.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/merval-project/merval/cmd/swarm
	go build -o bin/geth     github.com/merval-project/merval/cmd/geth
	go build -o bin/bootnode github.com/merval-project/merval/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/merval-project/merval/swarm/...

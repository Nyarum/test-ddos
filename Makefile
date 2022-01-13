CMD_CLIENT=cmd/client
CMD_SERVER=cmd/server

docker:
	docker build -f build/client.Dockerfile --tag=test-ddos-client .
	docker build -f build/server.Dockerfile --tag=test-ddos-server .

docker/create/network:
	docker network create local

docker/run/client:
	docker run -i --rm --network=local --name=client test-ddos-client -- addr=server:8999

docker/run/server:
	docker run -i --rm --network=local --name=server test-ddos-server

tests:
	go test ./...

build_client:
	go build -ldflags "-linkmode external -extldflags -static" -o client ./${CMD_CLIENT}

build_server:
	go build -ldflags "-linkmode external -extldflags -static" -o server ./${CMD_SERVER}

run_client: build_client
	./client addr=localhost:8999

run_server: build_server
	./server

.PHONY: tests build_client build_server
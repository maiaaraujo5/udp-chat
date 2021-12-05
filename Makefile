unit-test-server:
		go clean -testcache && go test -race ./internal/app/server...

unit-test-client:
		go clean -testcache && go test -race ./internal/app/client...

end-to-end-tests:
		go clean -testcache && go test -tags=integration ./e2e_test...

lint:
	golangci-lint run ./... --config ./build/golangci-lint/config.yml

docker-compose-run-redis:
		cd ./build/docker-compose && docker-compose up -d

build-server:
		go mod vendor
		go build -mod vendor -o ./dist/server/main ./cmd/chat/server/main.go
		cp ./config/server/*.yaml ./dist/server

build-client:
		go mod vendor
		go build -mod vendor -o ./dist/client/main ./cmd/chat/client/main.go

run-server-with-redis: docker-compose-run-redis
		CONF=./config/server/default.yaml,config/server/development.yaml go run cmd/chat/server/main.go

run-server:
		CONF=./config/server/default.yaml,config/server/development.yaml go run cmd/chat/server/main.go

run-client:
		CONF=./config/client/development.yaml go run cmd/chat/client/main.go

docker-run-server: build-server
		cp ./build/dockerfile/Dockerfile ./dist/server/
		cd ./dist/server/;docker build -t maiaaraujo5/udp-chat-server:latest .
		docker run -e CONF=./default.yaml,./development.yaml --rm -d --network host maiaaraujo5/udp-chat-server:latest

docker-run-server-with-redis: docker-compose-run-redis docker-run-server
unit-test-server:
		go clean -testcache && go test -race ./internal/app/server...

unit-test-client:
		go clean -testcache && go test -race ./internal/app/client...

end-to-end-tests:
		go clean -testcache && go test -tags=integration ./e2e_test...

docker-compose-run-redis:
		cd ./build && docker-compose up -d

build-server:
		go mod vendor
		go build -mod vendor -o ./dist/server/main ./cmd/chat/server/main.go

build-client:
		go mod vendor
		go build -mod vendor -o ./dist/client/main ./cmd/chat/client/main.go


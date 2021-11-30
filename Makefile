unit-test-server:
		go clean -testcache && go test -race ./internal/app/server...

unit-test-client:
		go clean -testcache && go test -race ./internal/app/client...

end-to-end-tests:
		go clean -testcache && go test -tags=integration ./e2e_test...





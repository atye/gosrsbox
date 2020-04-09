test:

	go test -v -cover -race ./...

cover-profile:

	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out
	rm -f coverage.out
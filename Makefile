OPENAPI_DIR=osrsboxapi/openapi

generate:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.0.0 generate \
    -i /local/${OPENAPI_DIR}/openapi.yaml \
    -g go \
	--package-name client \
    -o /local/${OPENAPI_DIR}/client
	rm -f \
		${OPENAPI_DIR}/client/.openapi-generator/FILES \
		${OPENAPI_DIR}/client/.gitignore \
		${OPENAPI_DIR}/client/.openapi-generator-ignore \
		${OPENAPI_DIR}/client/.travis.yml \
		${OPENAPI_DIR}/client/git_push.sh \
		${OPENAPI_DIR}/client/go.mod \
		${OPENAPI_DIR}/client/go.sum

test:
	go test -v -cover -race ./...

cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
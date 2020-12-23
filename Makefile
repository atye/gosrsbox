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

	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.0.0 generate \
    -i /local/${OPENAPI_DIR}/openapi.yaml \
    -g go-server \
	--package-name server \
    -o /local/${OPENAPI_DIR}/server
	rm -f \
		${OPENAPI_DIR}/server/.openapi-generator/FILES \
		${OPENAPI_DIR}/server/.gitignore \
		${OPENAPI_DIR}/server/.openapi-generator-ignore \
		${OPENAPI_DIR}/server/.travis.yml \
		${OPENAPI_DIR}/server/git_push.sh \
		${OPENAPI_DIR}/server/go.mod \
		${OPENAPI_DIR}/server/go.sum

test:
	go test -v -cover -race ./...

cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
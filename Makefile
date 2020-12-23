OPENAPI_DIR=pkg/openapi

generate:
	openapi-generator generate -i ${OPENAPI_DIR}/openapi.yaml -g go --additional-properties=generateInterfaces=true --package-name api \
		-o ${OPENAPI_DIR}/api
	rm -f \
		${OPENAPI_DIR}/api/.openapi-generator/FILES \
		${OPENAPI_DIR}/api/.gitignore \
		${OPENAPI_DIR}/api/.openapi-generator-ignore \
		${OPENAPI_DIR}/api/.travis.yml \
		${OPENAPI_DIR}/api/git_push.sh \
		${OPENAPI_DIR}/api/go.mod \
		${OPENAPI_DIR}/api/go.sum

test:
	go test -v -cover -race ./...

cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
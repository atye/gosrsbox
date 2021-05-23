OPENAPI_DIR=internal/openapi
 
 # openapi.yaml:
 # _id properties must be removed
 # id properties must be string, except for MonsterDrops
 # response meta page must be int
openapi:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.0.0 generate \
    -i /local/${OPENAPI_DIR}/openapi.yaml \
    -g go \
	--package-name api \
    -o /local/${OPENAPI_DIR}/api
	rm -f \
		${OPENAPI_DIR}/api/.openapi-generator/FILES \
		${OPENAPI_DIR}/api/.gitignore \
		${OPENAPI_DIR}/api/.openapi-generator-ignore \
		${OPENAPI_DIR}/api/.travis.yml \
		${OPENAPI_DIR}/api/git_push.sh \
		${OPENAPI_DIR}/api/go.mod \
		${OPENAPI_DIR}/api/go.sum

test:
	go test -count=1 -v -cover -race ./...

cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
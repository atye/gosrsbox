OPENAPI_DIR=openapi
 
 # openapi.json:
 # _id properties must be removed
 # id properties must be string, except for MonsterDrops
 # response meta page must be int
 .PHONY: openapi
openapi:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.1.1 generate \
    -i /local/${OPENAPI_DIR}/openapi.json \
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

.PHONY: test
test:
	go test -count=1 -v -cover -race ./...

.PHONY: cover-profile
cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
OPENAPI=internal/api
MODELS=models
 
 # openapi.json:
 # _id properties must be removed
 # id properties must be string, except for MonsterDrops
 # response meta page must be int

# apis,apiDocs=false,models,modelDocs=false,supportingFiles=client.go:configuration.go:utils.go
 .PHONY: openapi
openapi:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.1.1 generate \
    -i /local/openapi.json \
    -g go \
	--package-name testapi \
	--global-property apis,supportingFiles \
    -o /local/testapi
	rm -rf ${OPENAPI}/api


.PHONY: test
test:
	go test -count=1 -v -cover -race ./...


.PHONY: cover-profile
cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
OPENAPI=internal/client/openapi/api
 
 # openapi.json:
 # _id properties must be removed
 # id properties must be string, except for MonsterDrops
 # response meta page must be int

 .PHONY: openapi
openapi:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.1.1 generate \
    -i /local/openapi.json \
    -g go \
	--package-name api \
	--global-property apis=Item:Monster:Prayer,apiDocs=false,models,modelDocs=false,supportingFiles=client.go:configuration.go:utils.go \
    -o /local/${OPENAPI}
	rm -rf ${OPENAPI}/api

.PHONY:
mocks:
	mockgen -source internal/client/common/common.go -destination internal/client/common/mocks/mocks.go -package mocks


.PHONY: test
test:
	go test -count=1 -v -cover -race ./...


.PHONY: cover-profile
cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
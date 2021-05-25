OPENAPI=internal/openapi
MODELS=models
 
 # openapi.json:
 # _id properties must be removed
 # id properties must be string, except for MonsterDrops
 # response meta page must be int

 .PHONY: openapi
openapi:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.1.1 generate \
    -i /local/openapi.json \
    -g go \
	--package-name openapi \
	--global-property apis,apiDocs=false,models,modelDocs=false,supportingFiles=client.go:configuration.go:utils.go \
    -o /local/${OPENAPI}
	rm -rf ${OPENAPI}/api


.PHONY: models
models:
	docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli:v5.1.1 generate \
    -i /local/openapi.json \
    -g go \
	--package-name models \
	--global-property models="Item:Item_equipment:Item_weapon:Item_weapon_stances:Monster:Monster_drops:Prayer",,modelDocs=false,supportingFiles=utils.go \
    -o /local/models
	rm -rf ${MODELS}/api


.PHONY: test
test:
	go test -count=1 -v -cover -race ./...


.PHONY: cover-profile
cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
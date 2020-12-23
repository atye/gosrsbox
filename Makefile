OPENAPI_CLIENT_DIR=./osrsboxapi/api/internal/client/openapi

generate:
	curl -k https://api.osrsbox.com/api-docs | jq 'walk(if type == "object" then del(.properties._id) else . end)' > /tmp/osrsboxapi-openapi.json
	openapi-generator generate -i /tmp/osrsboxapi-openapi.json -g go --additional-properties=generateInterfaces=true --package-name openapi \
		-o ./osrsboxapi/api/internal/client/openapi
	rm -f \
		${OPENAPI_CLIENT_DIR}/.openapi-generator/FILES \
		${OPENAPI_CLIENT_DIR}/.gitignore \
		${OPENAPI_CLIENT_DIR}/.openapi-generator-ignore \
		${OPENAPI_CLIENT_DIR}/.travis.yml \
		${OPENAPI_CLIENT_DIR}/git_push.sh \
		${OPENAPI_CLIENT_DIR}/go.mod \
		${OPENAPI_CLIENT_DIR}/go.sum

test:
	go test -v -cover -race ./...

cover-profile:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm -f coverage.out
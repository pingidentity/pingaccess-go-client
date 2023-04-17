#!/usr/bin/env bash

set -x
set -e

# First remove any existing files that were generated. If we want to edit files post-generation, this section will need to be removed
# rm ./api_*.go
# rm ./model_*.go
# rm -r docs/

openapi-generator generate \
    -i ./api/openapi.yaml \
    -g go \
    --git-host github.com \
    --git-repo-id pingaccess-go-client \
    --git-user-id pingidentity \
    --type-mappings=integer=int64,number=float64 \
    --additional-properties=enumClassPrefix=true,packageName=configurationapi

rm -r test/

# Run any code generators
go mod tidy
go generate ./...

# Clean things up
go fmt ./...

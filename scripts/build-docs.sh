#!/bin/bash

set -e

build_template="${BUILD_TEMPLATE:-openapi.gen.yaml}"
source_list="${SOURCE_LIST:-api/docs/sources.list}"
openapi_file="${OPEN_API_FILE:-api/docs/docs.swagger.json}"
temp_file=$(mktemp)

# Generate the api definitions
# This is required as the proto build adds a bunch of default tags that are not helpful
buf generate --template "$build_template" $(sed 's/\(.*\)/--path api\/demo\/\1/' "$source_list")

# Remove top level tags as this will contain "HelloService" and other top levels tags that we don't want
jq -r 'del(.tags)' "$openapi_file" > "$temp_file"
cat "$temp_file" > "$openapi_file"

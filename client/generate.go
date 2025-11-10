// SPDX-License-Identifier: Apache-2.0

package client

//go:generate sh -c "set -e; if [ -n \"$OPENAPI_SPEC\" ]; then if [ ! -f \"$OPENAPI_SPEC\" ]; then echo \"missing OpenAPI spec: $OPENAPI_SPEC\" >&2; exit 1; fi; exec go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml \"$OPENAPI_SPEC\"; else if [ ! -f ../../landscape-openapi-docs/openapi/landscape_api.bundle.yaml ]; then echo \"missing OpenAPI spec: ../../landscape-openapi-docs/openapi/landscape_api.bundle.yaml\" >&2; exit 1; fi; exec go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ../../landscape-openapi-docs/openapi/landscape_api.bundle.yaml; fi"

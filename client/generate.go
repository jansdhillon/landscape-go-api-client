package client

//go:generate sh -c 'SPEC="${OPENAPI_SPEC:-../../landscape-openapi-docs/dist/openapi.bundle.yaml}" && go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml "$SPEC"'

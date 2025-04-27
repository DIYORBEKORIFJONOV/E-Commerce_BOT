# Makefile for E-Commerce_BoT/api_gateway

# Use “swag” by default for Swagger annotation parsing,
# but allow overriding via environment variable.
SWAGGER_CMD ?= swag

# Declare phony targets so Make doesn’t try to make files with these names.
.PHONY: all gen-product gen-order swagger-init run

# Generate Go code and gRPC stubs from your product protobuf
gen-product:
	@protoc -I pkg/protos pkg/protos/product/product.proto \
		--go_out=pkg/protos/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/protos/gen \
		--go-grpc_opt=paths=source_relative

# Generate Go code and gRPC stubs from your order protobuf
gen-order:
	@protoc -I pkg/protos pkg/protos/order/order.proto \
		--go_out=pkg/protos/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/protos/gen \
		--go-grpc_opt=paths=source_relative

# Initialize Swagger docs from your Swaggo comments
swagger-init:
	@$(SWAGGER_CMD) init -g internal/http/handler/order.go -o internal/app/docs

# Run the application (adjust as needed to point at your main package)
run:
	@go run cmd/main.go

# By default, build protos, generate Swagger, then run
all: gen-product gen-order swagger-init run

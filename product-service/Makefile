gen-product:
	@protoc -I pkg/protos pkg/protos/product/product.proto \
		--go_out=pkg/protos/gen \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/protos/gen \
		--go-grpc_opt=paths=source_relative
run:
	go run cmd/product/main.go
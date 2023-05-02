.PHONY: setup
setup:
	mkdir -p generated/key

.PHONY: generate
generate: key-service

.PHONY: key-service
key-service:
	protoc -I . \
		--go_out ./generated/key --go_opt paths=source_relative \
		--go-grpc_out ./generated/key --go-grpc_opt paths=source_relative \
		v1/key.proto

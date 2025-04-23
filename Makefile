# For generating gRPC and Go code from .proto files
# Change this to your service name, e.g. user, auth, payment, etc.
SERVICE ?= user

# Paths
PROTO_DIR := proto/$(SERVICE)
OUT_DIR := .

# Proto files to compile
PROTO_FILES := $(wildcard $(PROTO_DIR)/*.proto)

# Default target
.PHONY: proto
proto:
	@echo "Generating Go and gRPC code for service: $(SERVICE)"
	@protoc \
		--go_out=$(OUT_DIR) \
		--go-grpc_out=$(OUT_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_FILES)
	@echo "✅ Generation complete!"

# Clean generated files
.PHONY: clean
clean:
	@echo "Cleaning generated files for service: $(SERVICE)"
	@rm -rf $(OUT_DIR)/proto/$(SERVICE)
	@echo "🧹 Clean complete!"

# Define variables
PROTOC := protoc
PROTO_DIR := proto                              # Root directory where .proto files are stored
MODULE := github.com/oriastanjung/personal-va   # Go module name
OUT_DIR := .                                    # Output directory for generated files

# Find all .proto files in the proto directory and its subdirectories
PROTO_FILES := $(shell find $(PROTO_DIR) -name '*.proto')

# Rule to compile all proto files
.PHONY: proto
proto: ## Compile all protobuf files in the project
	$(PROTOC) -I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=module=$(MODULE) \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=module=$(MODULE) \
		$(PROTO_FILES)

# Clean generated .pb.go files
.PHONY: clean
clean: ## Clean generated protobuf files
	find $(PROTO_DIR) -name '*.pb.go' -delete
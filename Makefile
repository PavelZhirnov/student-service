API_PATH 		= api/studentService
PROTO_API_DIR 	= api/studentService
PROTO_OUT_DIR 	= pkg/studentServiceApi
PROTO_API_OUT_DIR = ${PROTO_OUT_DIR}
MOCKS_DIR = internal/repository
PROTO_API_VERSION = v2

.PHONY: gen-proto-ss
gen-proto-ss:
	mkdir -p ${PROTO_OUT_DIR}
	protoc \
		-I ${API_PATH} \
		-I third_party/googleapis \
		--include_imports \
		--go_out=$(PROTO_OUT_DIR) --go_opt=paths=source_relative \
        --go-grpc_out=$(PROTO_OUT_DIR)  --go-grpc_opt=paths=source_relative \
		--descriptor_set_out=$(PROTO_API_OUT_DIR)/api.pb \
		./${PROTO_API_DIR}/*.proto

.PHONY: go/lint
go/lint:
	golangci-lint run  --config=.golangci.yml --timeout=180s ./...


.PHONY: go-mock-install
go-mock-install:
	GO111MODULE=on go get github.com/golang/mock/mockgen@v1.5.0

.PHONY: mocks-generate
mocks-generate:
	@echo "generate-mocks"
	mockgen -package=repository -source=$(MOCKS_DIR)/interfaces.go -destination=$(MOCKS_DIR)/mock_gen.go
	@echo "successfully"

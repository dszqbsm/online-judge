SERVICE_NAME=onlinejudge
API_SERVICE_PATH=./api
RPC_SERVICE_PATH=./rpc

all: run_api

run_api:
	go run -v main.go api --config etc/api.yaml

build_api:
	goctl api go -api api/$(SERVICE_NAME).api -dir $(API_SERVICE_PATH) --style go_zero
	@if [ ! -f "./etc/api.yaml" ]; then mkdir -p ./etc && cp $(API_SERVICE_PATH)/etc/$(SERVICE_NAME).yaml ./etc/api.yaml; fi
	@rm -rf $(API_SERVICE_PATH)/etc

build_api_swagger:
	goctl api plugin -plugin goctl-swagger="swagger -filename api/$(SERVICE_NAME).json" -api api/$(SERVICE_NAME).api -dir .

run_rpc:
	go run -v main.go rpc --config etc/rpc.yaml

build_rpc:
	protoc $(RPC_SERVICE_PATH)/$(SERVICE_NAME).proto --go_out=$(RPC_SERVICE_PATH) --go-grpc_out=$(RPC_SERVICE_PATH) --validate_out="lang=go:$(RPC_SERVICE_PATH)"
	goctl rpc protoc $(RPC_SERVICE_PATH)/$(SERVICE_NAME).proto --go_out=$(RPC_SERVICE_PATH) --go-grpc_out=$(RPC_SERVICE_PATH) --zrpc_out=$(RPC_SERVICE_PATH) --style go_zero
	@if [ ! -f "./etc/rpc.yaml" ]; then mkdir -p ./etc && cp $(RPC_SERVICE_PATH)/etc/$(SERVICE_NAME).yaml ./etc/rpc.yaml; fi
	@rm -rf $(RPC_SERVICE_PATH)/etc

build_rpc_tool:
	# 下载对应平台的protobuf的版本：https://github.com/protocolbuffers/protobuf/releases
	# 将 include目录中的文件复制 $(GOPATH) 下
	@rm -rf protoc-gen-validate
	@git clone https://github.com/bufbuild/protoc-gen-validate.git
	@cd protoc-gen-validate && make build && cp -rf validate $(GOPATH)/include
	@rm -rf protoc-gen-validate

build_model:
	goctl model mysql ddl -src ./model/sql/onlinejudge.sql -dir ./model --style go_zero

test:
	go test -v ./...

bench:
	go test ./... -bench=.

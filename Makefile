GOPATH:=$(shell go env GOPATH)

############################## app ##############################
.PHONY: api_init
api_init:
	@goctl api go -api ./service/app/api/app.api --home=./goctl/template -dir ./service/app/api -style go_zero

.PHONY: rpc_init
rpc_init:
	@goctl rpc protoc ./service/app/rpc/app.proto --home=./goctl/template --go_out=./service/app/rpc --go-grpc_out=./service/app/rpc --zrpc_out=./service/app/rpc --style go_zero


.PHONY: run_api
run_api:
	@go run ./service/app/api/app.go -f ./service/app/api/etc/app.yaml

.PHONY: run_rpc
run_rpc:
	@go run ./service/app/rpc/app.go -f ./service/app/rpc/etc/app.yaml


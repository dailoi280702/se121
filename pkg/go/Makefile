generate_grpc_utils_pkg:
	protoc --proto_path=../../protos/ --go_out=./grpc/generated/utils --go_opt=paths=source_relative \
	--go-grpc_out=./grpc/generated/utils --go-grpc_opt=paths=source_relative \
	utils.proto;

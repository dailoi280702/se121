generate_search_service_pkg:
	protoc --proto_path=../protos/ --go_out=./pkg/search --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/search/ --go-grpc_opt=paths=source_relative \
	search.proto;


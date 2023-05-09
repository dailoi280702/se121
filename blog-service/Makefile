generate_user_service:
	protoc --proto_path=../protos/ --go_out=./pkg/blog --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/blog/ --go-grpc_opt=paths=source_relative \
	blog.proto;

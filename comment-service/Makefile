generate_user_service:
	protoc --proto_path=../protos/ --go_out=./pkg/comment --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/comment/ --go-grpc_opt=paths=source_relative \
	comment.proto;

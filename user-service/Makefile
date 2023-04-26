generate_user_service:
	protoc --proto_path=../protos/ --go_out=./userpb --go_opt=paths=source_relative \
	--go-grpc_out=./userpb --go-grpc_opt=paths=source_relative \
	user.proto;

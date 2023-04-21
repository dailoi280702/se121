generate_user_service:
	protoc --proto_path=../protos/ --go_out=./internal/service/user/ --go_opt=paths=source_relative \
	--go-grpc_out=./internal/service/user/ --go-grpc_opt=paths=source_relative \
	user.proto;

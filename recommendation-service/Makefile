generate_recommendation_service:
	protoc --proto_path=../protos/ --go_out=./pkg/recommendation --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/recommendation/ --go-grpc_opt=paths=source_relative \
	recommendation.proto;

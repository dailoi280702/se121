generate_car_service_pkg:
	protoc --proto_path=../protos/ --go_out=./pkg/car --go_opt=paths=source_relative \
	--go-grpc_out=./pkg/car/ --go-grpc_opt=paths=source_relative \
	car.proto;



.PHONY: proto
proto:
	protoc --proto_path=. --go-grpc_out=. --go_out=. proto/v1/authentication/authentication.proto
	protoc --proto_path=. --go-grpc_out=. --go_out=. proto/v1/user/user.proto
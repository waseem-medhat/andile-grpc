protogen:
	protoc \
		--go_out=internal/invoicer \
		--go_opt=paths=source_relative \
		--go-grpc_out=internal/invoicer \
		--go-grpc_opt=paths=source_relative \
		invoicer.proto

ui:
	grpcui -plaintext -proto invoicer.proto localhost:8080

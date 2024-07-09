.PHONY: greetings

greetings:
	protoc -Igreetings/proto \
 	--go_out=. \
 	--go_opt=module=github.com/andriesniemandt/go-grpc \
 	--go-grpc_out=. \
 	--go-grpc_opt=module=github.com/andriesniemandt/go-grpc \
 	greetings/proto/greetings.proto

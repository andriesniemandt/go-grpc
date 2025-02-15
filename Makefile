.PHONY: greetings
.PHONY: calculator
.PHONY: blog

greetings:
	protoc -Igreetings/proto \
 	--go_out=. \
 	--go_opt=module=github.com/andriesniemandt/go-grpc \
 	--go-grpc_out=. \
 	--go-grpc_opt=module=github.com/andriesniemandt/go-grpc \
 	greetings/proto/greetings.proto


calculator:
	protoc -Icalculator/proto \
	--go_out=. \
	--go_opt=module=github.com/andriesniemandt/go-grpc \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/andriesniemandt/go-grpc \
	calculator/proto/calculator.proto


blog:
	protoc -Iblog/proto \
	--go_out=. \
	--go_opt=module=github.com/andriesniemandt/go-grpc \
	--go-grpc_out=. \
	--go-grpc_opt=module=github.com/andriesniemandt/go-grpc \
	blog/proto/blog.proto
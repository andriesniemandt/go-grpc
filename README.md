# Go GRPC

Install the protobuf compiler plugins for golang:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Verify your `protoc` installation:

```bash
protoc --version
```

*libprotoc 3.21.12*

```bash
make greetings
```
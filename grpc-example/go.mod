module grpc-example

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.25.0
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.35.0

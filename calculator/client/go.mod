module github.com/ekkinox/go-grpc/calculator/client

go 1.18

replace github.com/ekkinox/go-grpc/calculator/proto => ../proto

require (
	github.com/ekkinox/go-grpc/calculator/proto v0.0.0-20220410125102-ded234fa5e4a
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/sys v0.0.0-20200323222414-85ca7c5b95cd // indirect
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)

package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/sdfwds4/test_go-micro_qps/proto"
	// "github.com/micro/plugins/v5/client/grpc"

	"go-micro.dev/v5"
	"go-micro.dev/v5/client"
)

func main() {
	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		// micro.Client(grpc.NewClient()),
	)

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := pb.NewGreeterService("helloworld", service.Client())
	cxt := context.Background()
	start := time.Now()

	var rsp *pb.Response
	var err error
	for {
		// Make request
		rsp, err = cl.Hello(cxt, &pb.Request{Name: "John"}, client.WithAddress("127.0.0.1:8081"))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
	}

	fmt.Println("duration:", time.Since(start))

	fmt.Println("rsp.Greeting:", rsp.Greeting)
}

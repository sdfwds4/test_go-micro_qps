package main

import (
	"context"
	"log"
	"sync/atomic"
	"time"

	pb "github.com/sdfwds4/test_go-micro_qps/proto"

	// "github.com/micro/plugins/v5/server/grpc"
	"go-micro.dev/v5"
)

var counter int64

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	rsp.Greeting = "Hello " + req.Name

	atomic.AddInt64(&counter, 1)

	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
		micro.Address("127.0.0.1:8081"),
		// micro.Server(grpc.NewServer()),
	)

	// optionally setup command line usage
	service.Init()

	// Register Handlers
	pb.RegisterGreeterHandler(service.Server(), new(Greeter))

	go func() {
		var t = time.Now().UnixNano() / 1e6
		for {
			select {
			case <-time.After(time.Second * 5):
				now := time.Now().UnixNano() / 1e6
				v := atomic.SwapInt64(&counter, 0)
				log.Print("count: ", float64(v)/float64((now-t)/1000), "/s")
				t = now
			}
		}
	}()

	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

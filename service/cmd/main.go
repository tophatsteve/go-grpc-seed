package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/tophatsteve/go-grpc-seed/service"
	"google.golang.org/grpc"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	lis, err := net.Listen("tcp", ":9191")

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	sample.RegisterSampleSvcServer(grpcServer, sample.NewService())

	go func() {
		grpcServer.Serve(lis)
	}()

	log.Printf("Running on port 9191")

	<-stop
	log.Printf("Stopping...")
	grpcServer.GracefulStop()
	log.Printf("Stopped")
}

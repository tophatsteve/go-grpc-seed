package main

import (
	"fmt"

	"github.com/tophatsteve/go-grpc-seed/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"localhost:9191",
		grpc.WithInsecure(),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := sample.NewSampleSvcClient(conn)
	ctx := context.Background()
	respGenerate, err := client.Generate(ctx, &sample.Empty{})

	fmt.Println(respGenerate.Value)
}

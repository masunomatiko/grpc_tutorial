package main

import (
	"context"
	"time"
	"log"
	"os"

	pb "github.com/masunomatiko/grpc_tutorial/echo/proto"
	"google.golang.org/grpc"


)

func init() {
	log.SetFlags(0)
	log.SetPrefix("[echo] ")
}

func main() {
	target := "localhost:50051"
	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %s", err)
	}
	defer conn.Close()

	client := pb.NewEchoServiceClient(conn)
	msg := os.Args[1]
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.Echo(ctx, &pb.EchoRequest{
		Message: msg
	})
	if err != nil {
		log.Println(err)
	}
	log.Println(r.GetMessage())

}

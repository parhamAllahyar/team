package grpc

// import (
// 	"context"
// 	"flag"
// 	pb "grpc-demo/proto"
// 	"log"
// 	"time"
// 

// 	"google.golang.org/grpc"
// 	"google.golang.org/grpc/credentials/insecure"

// )

// var (
// 	addr = flag.String("addr", "localhost:9000", "the address to connect to")

// 	name = flag.String("name", defaultName, "Name to greet")
// )

// func GrpcClient() {

// 	flag.Parse()

// 	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewNoteClient(conn)
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
// 	defer cancel()
// 	r, err := c.Say(ctx, &pb.CReq{Title: "aaaaa", Description: "bbbbbbbbbb"})
// 	if err != nil {
// 		log.Fatalf("could not : %v", err)
// 	}
// 	log.Printf("Greeting: %s", r)

// }

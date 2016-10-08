package main

import (
       "log"
       "os"

       "golang.org/x/net/context"
       "google.golang.org/grpc"
       "github.com/perigee/hubot/pb"
)

const (
      address = "localhost:50051"
)

func main() {
     conn, err := grpc.Dial(address, grpc.WithInsecure())
     if err != nil {
     	log.Fatalf("Client Failed: %v", err)
     }

     defer conn.Close()
     c := pb.NewInfoClient(conn)

     name := "hi"
     if len(os.Args) > 1 {
     	name = os.Args[1]
     }

     r, err := c.GetInfo(context.Background(), &pb.InfoRequest{Name: name})

     if err != nil {
     	log.Fatalf("Connection Failed: %v", err)
     }

     log.Printf("Info: %s, %s", r.Version, r.Endpoint)
}
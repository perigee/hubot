package main

import (
       //"fmt"
       "log"
       "net"

       "golang.org/x/net/context"
       "google.golang.org/grpc"
       pb "github.com/perigee/hubot/pb"
)

const (
     port = ":50051"
)

type server struct{}

func (s *server) GetInfo(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
     return &pb.InfoResponse{Version: "1.2", Endpoint: "hubot" }, nil
}

func main() {
     lis, err := net.Listen("tcp", port)
     
     if err != nil {
     	log.Fatalf("Failed: %v", err)
     }

     s := grpc.NewServer()
     pb.RegisterInfoServer(s, &server{})
     s.Serve(lis)
}

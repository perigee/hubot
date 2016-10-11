package main

import (
	//"fmt"
	"log"
	"net"
	"os/exec"

	pb "github.com/perigee/ubot/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetInfo(ctx context.Context, in *pb.InfoRequest) (*pb.InfoResponse, error) {
	log.Printf("Client: %s", in.Name)
	cmdName := "ls"
	cmdArgs := []string{"-al"}
	cmdOut, err := exec.Command(cmdName, cmdArgs...).Output()
	if err != nil {
		log.Printf("Error: %s", err)
		return &pb.InfoResponse{Version: "1.2", Endpoint: "ERROR"}, nil
	}

	outend := string(cmdOut)
	return &pb.InfoResponse{Version: "1.2", Endpoint: outend}, nil
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

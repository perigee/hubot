package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os/exec"

	pb "github.com/perigee/ubot/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	port = ":50051"
)

var (
	tls        = flag.Bool("tls", false, "Connection uses TLS if true, else plain TCP")
	certFile   = flag.String("cert_file", "testdata/server1.pem", "The TLS cert file")
	keyFile    = flag.String("key_file", "testdata/server1.key", "The TLS key file")
	jsonDBFile = flag.String("json_db_file", "testdata/route_guide_db.json", "A json file containing a list of features")
	port       = flag.Int("port", 50051, "The server port")
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

	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("Failed: %v", err)
	}

	var opts []grpc.ServerOption

	if *tls {

		creds, err := credentials.NewServerTLSFromFile(*cert_file, *key_file)

		if err != nil {
			log.Fatalf("Failed to generate creds: %v", err)
		}

		opts = []grpc.ServerOption{grpc.Creds(creds)}

	}

	s := grpc.NewServer(opts...)
	pb.RegisterInfoServer(s, &server{})
	s.Serve(lis)
}

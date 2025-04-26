package main

import (
	"fmt"
	"log"
	"net"

	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/config"
	"github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/connections"
	orderproduct "github.com/DIYORBEKORIFJONOV/E-Commerce_BOT.git/internal/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	c := config.Configuration()
	ls, err := net.Listen(c.User.Host, c.User.Port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	server := connections.NewService()
	orderproduct.RegisterOrderServiceServer(s, server)
	reflection.Register(s)
	fmt.Printf("server started on the port %s", c.User.Port)

	if err := s.Serve(ls); err != nil {
		log.Fatal(err)
	}
}

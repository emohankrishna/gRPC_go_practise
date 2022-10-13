package main

import (
	"context"
	"log"
	"net"

	"github.com/emohankrishna/gRPC_go_practise/helloworld"
	"github.com/emohankrishna/gRPC_go_practise/person"
	"google.golang.org/grpc"
)

type server struct {
	helloworld.UnimplementedGreeterServer
}

type server1 struct {
	person.UnimplementedGreetEmployeeServer
}

func (s *server) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	log.Printf("Received %v ", req.GetName())
	return &helloworld.HelloReply{Message: "Good Morning"}, nil
}

func (s *server1) GetEmployeeSendPerson(ctx context.Context, req *person.Employee) (*person.Person, error) {
	log.Println("**** Received Employee (Start)****")
	log.Printf("Employee Name %s", req.GetName())
	log.Printf("Employee Age %d", req.GetAge())
	log.Printf("Employee Sallary %f", req.GetSallary())
	log.Println("**** Received Employee (End)****")
	res := &person.Person{Firstname: "Mohan", Lastname: "Krishna"}
	return res, nil
}

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err.Error())
	}
	grpcServer := grpc.NewServer()

	helloworld.RegisterGreeterServer(grpcServer, &server{})
	person.RegisterGreetEmployeeServer(grpcServer, &server1{})
	log.Printf("server listening at %v", 9000)
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal("Failed to serve gRPC server over port 9000 :", err.Error())
	}

}

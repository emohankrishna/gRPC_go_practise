package main

import (
	"context"
	"log"
	"time"

	"github.com/emohankrishna/gRPC_go_practise/helloworld"
	"github.com/emohankrishna/gRPC_go_practise/person"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Client not connected ", err.Error())
	}
	c := helloworld.NewGreeterClient(conn)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: "Mohan"})
	if err != nil {
		log.Fatal("could not greet", err.Error())
	}
	log.Printf("Greets %s", r.GetMessage())
	p := person.NewGreetEmployeeClient(conn)
	res, err := p.GetEmployeeSendPerson(ctx, &person.Employee{
		Name:    "Hima Bindu",
		Age:     24,
		Sallary: 50000,
	})
	if err != nil {
		log.Fatal("could not greet", err.Error())
	}
	log.Println("**** Received Person (Start)****")
	log.Printf("Person firstname %s", res.GetFirstname())
	log.Printf("Person lastname %s", res.GetLastname())
	log.Println("**** Received Person (End)****")
}

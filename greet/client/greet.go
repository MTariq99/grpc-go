package main

import (
	"context"
	"fmt"
	pb "grpcgo/greet/proto"
	"log"
)

func doGreet(client pb.GreetServiceClient) {
	fmt.Println("doGreet function was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "tariq",
	})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)

}

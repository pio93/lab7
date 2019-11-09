package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/dat320/assignments/lab7/proto"
	"google.golang.org/grpc"
)

func main() {
	var (
		addr = flag.String("addr", "127.0.0.2:5000", "Client's address")
		sec  = flag.Int64("sec", 1, "Refresh rate in seconds")
		serv = flag.String("serv", "top10", "Choose service tha server will deliver")
	)

	flag.Parse()
	conn, err := grpc.Dial("127.0.0.1:4050", grpc.WithInsecure())
	log.Println("Dialing...")

	if err != nil {
		log.Fatalf("Couldn't connect to server %v", err)
	}

	client := proto.NewSubscriptionClient(conn)

	stream, err := client.Subscribe(context.Background(), &proto.SubscribeRequest{
		ClientAddr:  *addr,
		RefreshRate: *sec,
		Function:    *serv})

	if err != nil {
		log.Fatalf("Client: %v failed to create stream: %v ", client, err)
	}
	if *serv == "top10" {
		for {
			notification, err := stream.Recv()

			if err != nil {
				log.Fatalf("Failed to retrieve notification from server: %v", err)
				break
			}
			fmt.Println("\nTop10")
			for i, v := range notification.Channel {
				fmt.Printf("%d: %s : %d\n", i+1, v, notification.Viewers[i])
			}
		}
	} else if *serv == "durations" {
		for {
			fmt.Println("Durations")
			notification, err := stream.Recv()

			if err != nil {
				log.Fatalf("Failed to retriver notification from server: %v", err)
				break
			}
			if len(notification.Duration) == 0 {
				fmt.Println("No changes yet.")
			} else {
				for _, v := range notification.Duration {
					fmt.Println(v)
				}

			}
		}
	} else {
		fmt.Println("This service does not exist")
	}

}

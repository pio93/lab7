package main

import (
	"fmt"
	"log"
	"net"
	"sort"
	"time"

	"github.com/dat320/assignments/lab7/proto"
	"google.golang.org/grpc"
)

//Server is...
type Server struct {
	errChan chan error
}

//Subscribe is...
func (srv Server) Subscribe(subReq *proto.SubscribeRequest, stream proto.Subscription_SubscribeServer) error {
	seconds, _ := time.ParseDuration(fmt.Sprintf("%ds", subReq.RefreshRate))
	go func() {
		for {
			log.Printf("Sending to addrres : %s\n", subReq.ClientAddr)
			log.Printf("Waits for %d seconds\n", subReq.RefreshRate)
			time.Sleep(seconds)

			channels := ztore.ChannelsViewers()

			sort.SliceStable(channels, func(i, j int) bool {
				return channels[i].Viewers > channels[j].Viewers
			})

			topTen := make([]string, 0)

			for i, v := range channels {
				topTen = append(topTen, v.Channel)
				if i == 9 {
					break
				}
			}

			notification := proto.Notification{Channel: topTen}

			err := stream.Send(&notification)

			if err != nil {
				srv.errChan <- err
			}
		}
	}()
	log.Println("Ended stream")
	return <-srv.errChan
}

func startGRPC() {
	listener, err := net.Listen("tcp", ":4040")
	log.Println("Listeing to clients on port")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSubscriptionServer(grpcServer, Server{})
	grpcServer.Serve(listener)
}

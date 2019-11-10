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
	function := subReq.Function
	if function == "top10" {
		go func() {
			for {
				log.Printf("Sending to %s\n", subReq.ClientAddr)
				log.Printf("Waits for %d seconds\n", subReq.RefreshRate)
				time.Sleep(seconds)

				channels := ztore.ChannelsViewers()

				sort.SliceStable(channels, func(i, j int) bool {
					return channels[i].Viewers > channels[j].Viewers
				})

				topTen := make([]string, 0)
				ttViewers := make([]int64, 0)

				for i, v := range channels {
					topTen = append(topTen, v.Channel)
					ttViewers = append(ttViewers, int64(v.Viewers))
					if i == 9 {
						break
					}
				}

				notification := proto.Notification{Channel: topTen, Viewers: ttViewers}

				err := stream.Send(&notification)

				if err != nil {
					srv.errChan <- err
				}
			}
		}()
	} else if function == "durations" {
		go func() {
			for {
				log.Printf("Sending to %s\n", subReq.ClientAddr)
				log.Printf("Waits for %d seconds\n", subReq.RefreshRate)
				time.Sleep(seconds)

				stats := durtore.GetStats()
				result := make([]string, len(stats))

				copy(result, stats)

				durtore.ClearStats()

				notification := proto.Notification{Duration: result}
				err := stream.Send(&notification)

				if err != nil {
					srv.errChan <- err
				}

			}
		}()
	}
	log.Println("Ended stream")
	return <-srv.errChan
}

func startGRPC() {
	listener, err := net.Listen("tcp", ":4052")
	log.Println("Listeing to clients on port 4052")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterSubscriptionServer(grpcServer, Server{})
	grpcServer.Serve(listener)
}

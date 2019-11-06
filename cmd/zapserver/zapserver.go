// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"sort"
	"time"

	"github.com/dat320/assignments/lab7/zlog"
)

var chans chan *zlog.ChZap
var ztore zlog.ZapLogger

// runLab starts the server, sets up the zap storage, and runs the specified lab.
// Note that this function must not block.
func runLab(labNum, mcastAdr string) {
	start(mcastAdr)
	switch labNum {
	case "1.2":
		log.Println("For exercise 1.2, run: go test -v")
		os.Exit(0)
	case "1.3a", "1.3b", "1.4", "1.5":
		ztore = zlog.NewSimpleZapLogger()
	case "1.6", "2.x":
		ztore = zlog.NewViewersZapLogger()
	case "1.7":
		ztore = zlog.NewConcurrentZapLogger()
	}
	switch labNum {
	case "1.1":
		//TODO write code for dumping zap events to console
		chans = make(chan *zlog.ChZap)
		go dumpAll()
	case "1.3a":
		//TODO write code for recording and showing # of viewers on NRK1
		go showViewers("NRK1")
	case "1.3b":
		//TODO write code for recording and showing # of viewers on NRK1 and TV2 Norge
		go showViewers("TV2 Norge")
	case "1.4":
		//TODO write code for measurements
		go showViewersWithStats("NRK1")
	case "1.5", "1.6", "1.7":
		//TODO write code for top-10 list (generic for different data structures)
		go topTen()
	case "2.x":
		//TODO write code for publishing events to a subscriber client
		startGRPC()
	}
}

func dumpAll() {
	for v := range chans {
		fmt.Println(v.String())
	}
}

// start starts the zapserver event processing loop.
// Note that this function must not block.
func start(mcastAdr string) {
	log.Println("Starting ZapServer...")
	//TODO(student) write this method (5p)
	udpAddr, err := net.ResolveUDPAddr("udp", mcastAdr)

	log.Printf("Resolved addres: %s", mcastAdr)

	checkError(err, "Error resolving")

	conn, err := net.ListenMulticastUDP("udp", nil, udpAddr)

	log.Printf("Listening to address: %s", udpAddr)

	checkError(err, "Error Multicast")

	go handleEvent(conn)

}

func handleEvent(conn *net.UDPConn) {
	defer conn.Close()

	for {
		var buff [1024]byte
		n, _, err := conn.ReadFromUDP(buff[0:])

		if err != nil {
			log.Fatalf("Read error: %v", err)
			return
		}

		ch, _, _ := zlog.NewSTBEvent(string(buff[0:n]))

		if ch != nil {
			if ztore != nil {
				ztore.Add(*ch)
			} else {
				chans <- ch
			}

		}

	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Message: %s | Fatal error:  %s", msg, err.Error())
		os.Exit(1)
	}
}
func showViewers(chName string) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Viewers on %s : %d\n", chName, ztore.Viewers(chName))
	}

}

func showViewersWithStats(chName string) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Viewers on %s : %d\n", chName, ztore.Viewers(chName))

		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		fmt.Printf("Allocated memory: %v bytes\tTotal allocated memory: %v bytes\n", m.Alloc, m.TotalAlloc)
	}
}

func topTen() {
	for {
		time.Sleep(1 * time.Second)
		channels := ztore.ChannelsViewers()

		sort.SliceStable(channels, func(i, j int) bool {
			return channels[i].Viewers > channels[j].Viewers
		})

		fmt.Println("Top 10")
		for i, v := range channels {
			fmt.Printf("%d: %s\n", i+1, v.String())
			if i == 9 {
				break
			}
		}
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		fmt.Printf("Allocated memory: %v bytes\tTotal allocated memory: %v bytes\n", m.Alloc, m.TotalAlloc)
	}
}

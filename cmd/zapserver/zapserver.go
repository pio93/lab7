// +build !solution

// Leave an empty line above this comment.
package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/dat320/assignments/lab7/zlog"
)

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
		// ztore = zlog.NewViewersZapLogger()
	case "1.7":
		// ztore = zlog.NewConcurrentZapLogger()
	}
	switch labNum {
	case "1.1":
		//TODO write code for dumping zap events to console
		// go dumpAll()

	case "1.3a":
		//TODO write code for recording and showing # of viewers on NRK1
		go showerViewers("NRK1")
	case "1.3b":
		//TODO write code for recording and showing # of viewers on NRK1 and TV2 Norge
		go showerViewers("TV2 Norge")
	case "1.4":
		//TODO write code for measurements
	case "1.5", "1.6", "1.7":
		//TODO write code for top-10 list (generic for different data structures)
	case "2.x":
		//TODO write code for publishing events to a subscriber client
	}
}

// start starts the zapserver event processing loop.
// Note that this function must not block.
func start(mcastAdr string) {
	log.Println("Starting ZapServer...")
	//TODO(student) write this method (5p)
	udpAddr, err := net.ResolveUDPAddr("udp", mcastAdr)
	checkError(err, "Error resolving")
	conn, err := net.ListenMulticastUDP("udp", nil, udpAddr)
	checkError(err, "Error Multicast")
	defer conn.Close()

	for {
		buff := make([]byte, 1024)

		n, err := conn.Read(buff[0:])

		if err != nil {
			log.Fatalf("Read error: %v", err)
			return
		}

		fmt.Println(string(buff[0:n]))
	}
}

func checkError(err error, msg string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Message: %s | Fatal error:  %s", msg, err.Error())
		os.Exit(1)
	}
}
func showerViewers(chName string) {
	for {
		time.Sleep(1 * time.Second)
		fmt.Printf("Viewers on %s : %d", chName, ztore.Viewers(chName))
	}
}

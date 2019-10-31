// Zap Collection Server
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
)

func main() {
	var (
		maddr      = flag.String("mcast", "224.0.1.130:10000", "multicast ip:port")
		lab        = flag.String("lab", "1.1", "which lab exercise to run")
		memprofile = flag.String("memprofile", "", "write memory profile to this file")
	)
	flag.Parse()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Kill, os.Interrupt)
	runLab(*lab, *maddr)

	// wait for CTRL-C or some other kill signal
	s := <-signalChan
	fmt.Printf("\nServer stopping on %s signal\n", s)
	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
		fmt.Println("Saved memory profile")
		fmt.Println("Analyze with: go tool pprof zapserver", *memprofile)
	}
}

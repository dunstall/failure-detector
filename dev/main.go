package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/dunstall/failure-detector/monitor"
)

func args() (monitor.NodeID, uint16) {
	id := flag.Int("id", -1, "ID of this node")
	port := flag.Int("port", -1, "port to listen for and send heartbeats")
	flag.Parse()
	if int(uint32(*port)) != *port {
		fmt.Println("ID is required")
		os.Exit(1)
	}
	if int(uint16(*port)) != *port {
		fmt.Println("port is required")
		os.Exit(1)
	}

	return monitor.NodeID(uint32(*id)), uint16(*port)
}

func main() {
	fmt.Println("starting failure detector...")
	id, port := args()
	m, err := monitor.NewUDPMonitor(id, "255.255.255.255", port, time.Second)
	if err != nil {
		fmt.Println(err)
	}
	defer m.Close()

	for {
		select {
		case id := <-m.Heartbeats():
			fmt.Printf("heartbeat %#v\n", id)
		}
	}
}

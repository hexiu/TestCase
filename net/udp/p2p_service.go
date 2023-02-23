package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	Run()
}

func Run() {

	SAddr, _ := net.ResolveUDPAddr("udp4", "0.0.0.0:8888")

	ls, err := net.ListenUDP("udp4", SAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	var peers = make([]string, 0)
	var b = make([]byte, 1024)
	go func() {
		for {
			n, s, err := ls.ReadFromUDP(b)
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("resv:", string(b[:n]))
			peer := s.String()
			peers = append(peers, peer)
		}
	}()

	for {
		if len(peers) == 2 {
			rAddr, _ := net.ResolveUDPAddr("udp4", peers[1])
			ls.WriteToUDP([]byte(peers[0]), rAddr)
			rAddr, _ = net.ResolveUDPAddr("udp4", peers[0])
			ls.WriteToUDP([]byte(peers[1]), rAddr)
			return
		}
		time.Sleep(1 * time.Second)
		fmt.Println(time.Now().String(), peers)
	}

}

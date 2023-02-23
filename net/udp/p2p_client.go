package main

import (
	"fmt"
	"net"

	"log"
)

func main() {

	Run()
}

func Run() {
	UdpClient()
}

func UdpClient() {

	udpSrv, err := net.ResolveUDPAddr("udp4", "0.0.0.0:9001")
	if err != nil {
		log.Fatal("Unable to resolve: ", err)
		return
	}

	rddr, _ := net.ResolveUDPAddr("udp4", "0.0.0.0:8888")
	// 请求 srv
	conn, err := net.DialUDP("udp4", udpSrv, rddr)
	if err != nil {
		log.Fatal("Unable to DialUDP: ", err)
		return
	}
	_, err = conn.Write([]byte("hi"))
	if err != nil {
		log.Fatal("Unable to DialUDP: ", err)
		return
	}

	b := make([]byte, 1000)
	n, _, err := conn.ReadFromUDP(b)
	if err != nil {
		log.Fatal("Unable to DialUDP: ", err)
		return
	}
	perrAddr, _ := net.ResolveUDPAddr("udp4", string(b[:n]))
	fmt.Println("msg: ", perrAddr)
	conn.Close()

	// 请求 远程的 udp 接口
	rConn, err := net.DialUDP("udp4", udpSrv, perrAddr)
	if err != nil {
		log.Fatal("Unable to DialUDP: ", err)
		return
	}

	rConn.Write([]byte("hello"))

	go func() {
		for {
			b := make([]byte, 1024)
			n, _, err := rConn.ReadFromUDP(b)
			if err != nil {
				log.Fatal("Unable to DialUDP: ", err)
				return
			}
			fmt.Println("resv:", string(b[:n]))
		}

	}()

	for {
		var str string
		fmt.Scan(&str)
		rConn.Write([]byte(str))
	}

}

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
	Service()
}

func Service() {

	// laddr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:9011")
	raddr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:6666")

	tSrv, err := net.ListenTCP("tcp4", raddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	addlist := make(map[string]net.Conn, 0)
	addlist1 := make([]string, 0)
	for {
		cConn, err := tSrv.Accept()
		if err != nil {
			log.Fatal(err)
			return
		}
		addr := cConn.RemoteAddr().String()
		addlist[addr] = cConn
		fmt.Println(addlist)

		if len(addlist) == 2 {
			for k := range addlist {
				addlist1 = append(addlist1, k)
			}
			addlist[addlist1[0]].Write([]byte(addlist1[1]))
			addlist[addlist1[1]].Write([]byte(addlist1[0]))
			addlist[addlist1[0]].Close()
			addlist[addlist1[0]].Close()
			addlist1 = []string{}
			addlist = make(map[string]net.Conn, 0)
		}
		time.Sleep(time.Second)
	}

}

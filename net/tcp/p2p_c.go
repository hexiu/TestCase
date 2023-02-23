package main

import (
	"fmt"
	"log"
	"net"
	"time"

	reuseport "github.com/kavu/go_reuseport"
)

func main() {
	fmt.Println("\347\224\263\350\257\267\344\272\206\350\277\236\351\272\246")
	Run()
}

func Run() {
	Client()
}

func Client() {

	laddr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:9012")
	raddr, _ := net.ResolveTCPAddr("tcp4", "h2.jaxiu.cn:6666")

	tConn, err := net.DialTCP("tcp4", laddr, raddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	tConn.Write([]byte("hi"))

	b := make([]byte, 1024)
	n, err := tConn.Read(b)
	if err != nil {
		log.Fatal(err)
		return
	}
	ret := string(b[:n])
	fmt.Println(ret)
	fmt.Println(map[int64]int64{})

	// err = tConn.Close()
	fmt.Println(tConn, err)

	time.Sleep(time.Second)
	p2pAddr, _ := net.ResolveTCPAddr("tcp4", ret)
	var p2pConn *net.TCPConn
	go func() {
		for {
			fmt.Println("test")
			p2pConn, err = net.DialTCP("tcp4", nil, p2pAddr)
			if err != nil {
				fmt.Println("Err:", err)
				time.Sleep(time.Second)
				continue
			} else {
				fmt.Println("Break", err)
				break
			}
		}
		fmt.Println(err, "link ok", p2pConn.RemoteAddr().String())
	}()
	go func() {
		time.Sleep(time.Second)
		Conn, err := reuseport.Listen("tcp4", laddr.String())
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("listen ok ", Conn.Addr().String())

		c, err := Conn.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		for {
			b := make([]byte, 1024)
			n, err = c.Read(b)
			if err != nil {
				fmt.Println(err)
				return
			}
			ret := string(b[:n])
			fmt.Println(ret)
		}
	}()

	for {
		var str string
		fmt.Scan(&str)
		p2pConn.Write([]byte(str))
	}

}

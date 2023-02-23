package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	Run()
}

func Run() {
	fmt.Println("\346\236\204\345\273\272\345\244\261\350\264\245\357\274\201")
	Client()
}

func Client() {

	// laddr, _ := net.ResolveTCPAddr("tcp4", "0.0.0.0:9011")
	raddr, _ := net.ResolveTCPAddr("tcp4", "h2.jaxiu.cn:6666")

	tConn, err := net.DialTCP("tcp4", nil, raddr)
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

	// time.Sleep(time.Second)
	// p2pAddr, _ := net.ResolveTCPAddr("tcp4", ret)

	var p2pConn *net.TCPConn = tConn
	// for {
	// 	p2pConn, err = net.DialTCP("tcp4", nil, p2pAddr)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		time.Sleep(time.Second)
	// 		p2pConn = tConn
	// 		break
	// 	} else {
	// 		tConn.Close()
	// 		break
	// 	}
	// }
	// fmt.Println(err, "link ok", p2pConn.RemoteAddr().String())
	go func() {
		for {
			b := make([]byte, 1024)
			n, err := p2pConn.Read(b)
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
		_, err = p2pConn.Write([]byte(str))
		if err != nil {
			fmt.Println(err)
		}
	}

}

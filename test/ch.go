package test

import (
	"fmt"
	"time"
)

var chinfo = make(chan int, 10)
var close chan int

func Channel() {
	PublicCh()
	ConsumeCh()
	time.Sleep(time.Second)
}

func PublicCh() {
	go func() {
		for i := 0; i < 100; i++ {
			if i == 89 {
				close <- 1
			} else {
				chinfo <- i
			}
		}
	}()
}

func ConsumeCh() {
	go func() {
		for {
			select {
			case j := <-chinfo:
				fmt.Println(j)
			case <-close:
				fmt.Println("close")
				time.Sleep(100 * time.Millisecond)
				break
			}
		}
	}()

}

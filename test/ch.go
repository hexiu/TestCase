package test

import (
	"fmt"
	"time"
)

var chinfo = make(chan int, 10)
var close = make(chan int, 1)

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
			case k := <-close:
				fmt.Println("close", k)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()

}

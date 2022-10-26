package test

import (
	"fmt"
	"time"
)

var chinfo = make(chan int, 10)

func Channel() {
	PublicCh()
	ConsumeCh()
	time.Sleep(time.Second)
}

func PublicCh() {
	go func() {
		for i := 0; i < 100; i++ {
			chinfo <- i
		}
	}()
}

func ConsumeCh() {
	go func() {
		for i := 0; i < 100; i++ {
			select {
			case j := <-chinfo:
				fmt.Println(j)
			}
		}
	}()

}

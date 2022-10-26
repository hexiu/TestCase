package main

import (
	"fmt"
	"time"

	"github.com/looplab/fsm"
)

func main() {
	fsmTest()
}

func getFsm() *fsm.FSM {
	return fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "open"},
			{Name: "close", Src: []string{"open"}, Dst: "closed"},
			{Name: "mid", Src: []string{"open"}, Dst: "mid"},
		},

		fsm.Callbacks{
			"before_open": func(event *fsm.Event) {
				fmt.Println("before open callback ing........", event.Src, event.Dst, "event:", event.Event)
			},
			"open": func(event *fsm.Event) {
				fmt.Println("open callback ing ing........", event.Src, event.Dst, "event:", event.Event)
			},
			"after_open": func(event *fsm.Event) {
				fmt.Println("can:", event.FSM.Can("mid"), event.FSM.Is("open"))
				event.FSM.Is("open")
				fmt.Println("after_open callback: ", event.FSM.Current())
			},
			"close": func(event *fsm.Event) {
				fmt.Println("closecallback cancel open ing ing........", event.Src, event.Dst, "event:", event.Event)
			},
		},
	)
}

func fsmTest() {
	fsme := getFsm()
	
	go func() {
		fsme := getFsm()
		fsme.SetState("open")
		fmt.Println(fsme.Current())
		err := fsme.Event("close")
		fmt.Println("err:", err)
		fmt.Println(fsme.Current())
	}()
	go func() {
		fsme := getFsm()
		fsme.SetState("open")
		fmt.Println(fsme.Current())
		err := fsme.Event("close")
		fmt.Println("err:", err)
		fmt.Println(fsme.Current())
		fsme.SetState("open")

	}()
	fsme.SetState("open")
	fmt.Println(fsme.Current())
	err := fsme.Event("close")
	fmt.Println("err:", err)
	fmt.Println(fsme.Current())
	fsme.Event("open")
	fmt.Println(fsme.Current())

	time.Sleep(1 * time.Second)
	// for i := 0; i < 100; i++ {
	// 	go func() {

	// 		fmt.Println(fsme.Current())

	// 		err := fsme.Event("open")
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		fmt.Println(fsme.Current())

	// 		err = fsme.Event("close")
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}

	// 		fmt.Println(fsme.Current())

	// 	}()
	// }
	// time.Sleep(2 * time.Second)
}

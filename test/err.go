package test

import (
	"errors"
	"fmt"
)

type Aerr struct {
}

func (a *Aerr) Error() string {
	return "aerr"
}

// func TestFunc(ctx context.Context) {
// 	for i := 0; i < 1; i++ {
// 		go func1(ctx)
// 	}
// 	time.Sleep(time.Second)
// }

func AA3() {
	fmt.Println("aa3: ", AA2())
}

func AA2() (err error) {
	i, err := AA1()
	if err != nil {
		return
	}
	fmt.Println(i, err)
	return
}

func AA1() (i int64, err error) {
	return 1, errors.New("1")
}

// func func1(ctx context.Context) (err error) {
// 	wg, _ := errgroup.WithContext(ctx)
// 	for i := 0; i < 1000; i++ {
// 		j := i
// 		wg.Go(func() error {
// 			var err error
// 			if j%2 == 1 {
// 				err = &Aerr{}
// 			} else {
// 				err = fmt.Errorf("%v,,,", j)
// 			}
// 			fmt.Printf("%v,,,,%v\n", j, err)
// 			return err
// 		})
// 	}
// 	err = wg.Wait()
// 	fmt.Println(err)
// 	return
// }

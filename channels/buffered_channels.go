package main

import (
	"fmt"
)

func tree() {

}

func main() {

	//make a buffered channel of buffer size 3
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
		}
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

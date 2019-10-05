package main

import (
	"fmt"
)

func timestable(n int, ch chan string) {
	for i := 1; i <= 12; i++ {
		ch <- fmt.Sprintf("%d x %d = %d", n, i, (i * n))
	}
	close(ch)

}

func times12(ch chan string) {
	for n := 1; n <= 12; n++ {
		ch1 := make(chan string)
		go timestable(n, ch1)
		done := false
		for {
			select {
			case r, ok := <-ch1:
				if !ok {
					done = true
					break
				}
				ch <- r
			}

			if done {
				break
			}
		}
	}
	close(ch)
}
func main() {

	ch1 := make(chan string)

	ch2 := make(chan string)

	chAll := make(chan string)

	go timestable(2, ch1)
	go timestable(3, ch2)
	go times12(chAll)

	done1 := false
	done2 := false
	doneall := false

	for {
		select {
		case r1, ok := <-ch1:
			if !ok {
				done1 = true
				break
			}

			fmt.Println(r1)
		case r2, ok := <-ch2:
			if !ok {
				done2 = true
				break
			}
			fmt.Println(r2)

		case r, ok := <-chAll:

			if !ok {
				doneall = true
				break
			}
			fmt.Println(r)
		}

		if done1 && done2 && doneall {
			break
		}

	}
}

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {	
	go supervisor(count, 3)

	select {}
}

func Upcase(name string) string {
	return strings.ToUpper(name)
}

func Lowcase(name string) string {
	return strings.ToLower(name)
}

func SumArray(nums []int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
func Sum(x int, y int) int {
	return x + y
}

func count(to int) {
	for i := 0; i <= to; i++ {
		fmt.Printf("i=%d\n", i)
		time.Sleep(time.Second)
	}

	panic("pretend that something broke")
}

// Erlang-like supervision in go
func supervisor(fn func(int), args int) {
	for {
		func() {
			defer func() {
				if err := recover(); err != nil {
					fmt.Println("panic recovered")
				}
			}()
			fn(args)
		}()
		fmt.Println("restarting process")
		time.Sleep(time.Second)
	}
}

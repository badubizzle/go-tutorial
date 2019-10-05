package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func timestable(n int) {
	defer wg.Done()
	for i := 1; i <= 12; i++ {

		fmt.Println(i * n)
	}
}

func main() {

	go func() {

	}()
}

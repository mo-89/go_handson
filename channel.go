package main

import (
	"fmt"
	"time"
)

func total(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
	// c <- t
}

func main() {
	c := make(chan int)
	go total(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
	// go total(100, c)
	// fmt.Println("total: ", <-c)

	// go total(1000, c)
	// go total(100, c)
	// go total(10, c)
	// x, y, z := <-c, <-c, <-c
	// fmt.Println(x, y, z)
}

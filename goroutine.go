package main

import (
	"fmt"
	"time"
)

func Goroutine1() {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine 1")
		time.Sleep(100 * time.Millisecond)
	}
}

func Goroutine2() {
	for i := 0; i < 5; i++ {
		fmt.Println("goroutine2")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go Goroutine1()
	go Goroutine2()
	time.Sleep(1500 * time.Millisecond)
}

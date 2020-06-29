package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			go worker(channel, i)
		}
	}()

	for i := 0; i < 100; i++ {
		channel <- i
	}
}

func worker(channel chan int, worker int) {
	for i := range channel {
		fmt.Println("The number ", i, " was processed for worker ", worker)
		time.Sleep(time.Duration(rand.Intn(4)) * time.Second)
	}
}

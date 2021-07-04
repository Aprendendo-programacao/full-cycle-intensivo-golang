package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan int)

	go func() {
		for i := 0; i <= 5; i++ {
			go worker(channel)
		}
	}()

	// o "channel" ficará ocupado (não receberá um novo valor) até que um dos 5 workers sejam liberados
	for i := 0; i <= 100; i++ {
		channel <- i
	}

}

func worker(channel chan int) {
	for i := range channel {
		fmt.Println(i)
		time.Sleep(time.Second * 3)
	}
}
package main

import (
    "fmt"
    "time"
)

func work(done <-chan interface{}, msg <-chan string) <-chan interface{} {
	terminated := make(chan interface{})
	go func() {
		defer fmt.Println("exit work")
		defer close(terminated)

		for {
			select {
			case m := <-msg:
				fmt.Println(m)
			case <-done:
				return
			}
		}
	}()
	return terminated
}

func main() {
	done := make(chan interface{})
	msg := make(chan string)
	terminated := work(done, msg)

	msg <- "hello"

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("stopping goroutine")
		close(done)
	}()

    go func() {
        time.Sleep(1 * time.Second)
        msg <- "world"
    }()

	<-terminated

	fmt.Println("done")
}

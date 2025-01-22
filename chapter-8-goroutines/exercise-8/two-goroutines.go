package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	exit := make(chan struct{})

	var counter int

	go func() {
		for {
			select {
			case <-ch1:
				counter++
				ch2 <- struct{}{}
			case <-exit:
				return
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ch2:
				counter++
				ch1 <- struct{}{}
			case <-exit:
				return
			}
		}
	}()

	go func() {
		ch1 <- struct{}{}
		time.Sleep(time.Second)
		close(exit)
	}()

	<-exit
	fmt.Println(counter)
}

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ch <- "ping"
			case <-done:
				return
			}
		}
	}()

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				ch <- "pong"
			case <-done:
				return
			}
		}
	}()

	go func() {
		timer := time.NewTimer(time.Second * 5)
		defer timer.Stop()

		for {
			select {
			case m := <-ch:
				fmt.Println(m)
			case <-timer.C:
				done <- struct{}{}
			}
		}
	}()

	<-done
}

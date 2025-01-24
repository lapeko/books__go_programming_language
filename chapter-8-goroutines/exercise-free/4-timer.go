package main

import (
	"fmt"
	"time"
)

func main() {
	tick := make(chan string)
	done := make(chan struct{})

	go func() {
		ticker := time.NewTicker(time.Second * 2)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				tick <- t.Format("15:04:05")
			}
		}
	}()

	go func() {
		timer := time.NewTimer(time.Second * 10)
		defer timer.Stop()

		for {
			select {
			case t := <-tick:
				fmt.Println(t)
			case <-timer.C:
				close(done)
			}
		}
	}()

	<-done
}

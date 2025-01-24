package main

import (
	"errors"
	"log"
	"math/rand"
)

func main() {
	const (
		processGenLimit = 100
		attemptsOnFail  = 3
	)
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < processGenLimit; i++ {
			ch <- genTask()
		}
		close(ch)
	}()

	go func() {
		for id := range ch {
			attempt := 0
			for {
				if attempt == attemptsOnFail {
					log.Printf("Task %d failed after %d attempts\n", id, attempt)
					break
				}
				attempt++
				if err := riskyTask(id); err == nil {
					log.Printf("Task %d completed successfully on attempt %d\n", id, attempt)
					break
				}
			}
		}
		close(done)
	}()

	<-done
}

func genTask() int {
	return rand.Intn(1000)
}

func riskyTask(taskId int) error {
	if rand.Intn(2) == 0 {
		return errors.New("error")
	}
	return nil
}

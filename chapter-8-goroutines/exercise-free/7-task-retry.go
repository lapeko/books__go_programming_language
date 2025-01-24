package main

import (
	"errors"
	"log"
	"math/rand"
	"sync"
)

const (
	processGenLimit = 100
	attemptsOnFail  = 3
)

type retryResponse struct {
	ok       bool
	attempts int
	id       int
}

func main() {
	ch := make(chan int)
	done := make(chan struct{})

	go func() {
		for i := 0; i < processGenLimit; i++ {
			ch <- genTask()
		}
		close(ch)
	}()

	go func() {
		wg := &sync.WaitGroup{}
		res := make(chan retryResponse)

		for id := range ch {
			wg.Add(1)
			go runProcessWithRetry(id, res)
		}

		go func() {
			for r := range res {
				if r.ok {
					log.Printf("Task %d completed successfully on attempt %d\n", r.id, r.attempts)
				} else {
					log.Printf("Task %d failed after %d attempts\n", r.id, r.attempts)
				}
				wg.Done()
			}
		}()

		wg.Wait()
		close(res)
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

func runProcessWithRetry(id int, res chan<- retryResponse) {
	attempt := 0
	for {
		if attempt == attemptsOnFail {
			res <- retryResponse{ok: false, attempts: attempt, id: id}
			return
		}
		attempt++
		if err := riskyTask(id); err == nil {
			res <- retryResponse{ok: true, attempts: attempt, id: id}
			return
		}
	}
}

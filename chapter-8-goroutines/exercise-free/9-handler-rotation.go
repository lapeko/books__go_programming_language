package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

type task func() int

func main() {
	const taskHandlersCount = 3
	const numberOfTasks = 50

	wg := &sync.WaitGroup{}
	tasks := make(chan task, 20)

	for i := 1; i <= taskHandlersCount; i++ {
		go taskHandler(tasks, wg, i)
	}

	for i := 1; i <= numberOfTasks; i++ {
		wg.Add(1)
		tasks <- func() int {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
			return i
		}
	}
	close(tasks)

	wg.Wait()
}

func taskHandler(tc chan task, wg *sync.WaitGroup, handlerNum int) {
	for t := range tc {
		log.Printf("Handler %d started a new process...", handlerNum)
		log.Printf("Handler %d complited a process with ID: %d", handlerNum, t())
		wg.Done()
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	logChan := make(chan string, 4)
	shutdown := make(chan struct{})

	go func() {
		defer close(logChan)
		for {
			select {
			case logChan <- genLogWithDelay():
			case <-shutdown:
				return
			}
		}
	}()

	go func() {
		logBuf := make([]string, 0, 4)
		timer := time.NewTimer(time.Second * 5)
		defer timer.Stop()

		defer func() {
			logAndCleanBuff(&logBuf)
			close(shutdown)
		}()

		for {
			select {
			case log := <-logChan:
				logBuf = append(logBuf, log)
				if len(logBuf) >= 4 {
					logAndCleanBuff(&logBuf)
				}
			case <-timer.C:
				return
			}
		}
	}()

	<-shutdown
}

var logCounter int

func genLogWithDelay() string {
	logCounter++
	randDuration := time.Millisecond * time.Duration(rand.Intn(500))
	time.Sleep(randDuration)
	return fmt.Sprintf("Log message %d", logCounter)
}

func logAndCleanBuff(buff *[]string) {
	if len(*buff) == 0 {
		return
	}
	fmt.Println(*buff)
	*buff = (*buff)[:0]
}

package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	imgs := []string{"tiger", "dog", "flower"}
	rendered, err := makeThumbnailsInParallel(imgs)
	fmt.Printf("rendered images: %v\nerror: %v\n", rendered, err)
	fmt.Println("Parallel work Done")
}

func makeThumbnailsInParallel(imgs []string) ([]string, error) {
	type item struct {
		ok  string
		err error
	}

	ch := make(chan item)
	exit := make(chan struct{})
	var madeImgs []string

	for _, img := range imgs {
		go func(img string) {
			var it item
			it.ok, it.err = makeThumbnail(img)
			select {
			case ch <- it:
			case <-exit:
				return
			}
		}(img)
	}

	for range imgs {
		if i := <-ch; i.err != nil {
			close(exit)
			return madeImgs, i.err
		} else {
			madeImgs = append(madeImgs, i.ok)
		}
	}

	close(exit)
	return madeImgs, nil
}

var counter = 0

func makeThumbnail(path string) (string, error) {
	fmt.Println("makeThumbnail start")
	time.Sleep(time.Second)

	counter++

	if counter == 3 {
		time.Sleep(time.Millisecond * 100)
		return "", errors.New("unexpected error occurred")
	}

	fmt.Println("makeThumbnail finish")
	return fmt.Sprintf("%s-thumb", path), nil
}

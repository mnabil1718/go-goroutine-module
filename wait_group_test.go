package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func PrintNumber(num int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	waitGroup.Add(1)
	time.Sleep(1 * time.Second)
	fmt.Println(num)
}

func TestWaitGroup(t *testing.T) {

	waitGroup := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go PrintNumber(i, waitGroup)
	}

	waitGroup.Wait()

	fmt.Println("All processes done")
}

package main

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func PutDataIntoChannel(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Hello"
	fmt.Println("CHANNEL: Put data success")
}

func InputDataOnly(channel chan<- string) {
	channel <- "Hello"
	fmt.Println("CHANNEL: input data success")
}

func OutputDataOnly(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
	fmt.Println("CHANNEL: output data success")
}

func RangedChannel(channel chan string) {
	defer close(channel)
	for i := 0; i < 10; i++ {
		channel <- "Data number " + strconv.Itoa(i)
	}
}

func TestChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go PutDataIntoChannel(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannelDirection(t *testing.T) {
	channel := make(chan string)
	defer close(channel)
	go InputDataOnly(channel)
	go OutputDataOnly(channel)

	time.Sleep(5 * time.Second)
}

func TestRangedChannel(t *testing.T) {
	channel := make(chan string)

	go RangedChannel(channel)

	for data := range channel {
		fmt.Println(data)
	}

	time.Sleep(5 * time.Second)
}

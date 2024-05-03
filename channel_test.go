package main

import (
	"fmt"
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

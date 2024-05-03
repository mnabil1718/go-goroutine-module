package main

import (
	"fmt"
	"testing"
	"time"
)

func SayHello() {
	fmt.Println("Hello")
}

func TestSayHello(t *testing.T) {

	go SayHello()

	fmt.Println("I'm First")
	time.Sleep(1 * time.Second)
}

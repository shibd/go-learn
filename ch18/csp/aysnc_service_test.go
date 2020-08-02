package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	otherTask()
}

func AsyncService() chan string {
	retch := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("return result.")
		retch <- ret
		fmt.Println("service exited.")
	}()
	return retch
}

func TestAsyncService(t *testing.T) {
	retch := AsyncService()
	otherTask()
	fmt.Println(<-retch)
	time.Sleep(time.Second * 1)
}

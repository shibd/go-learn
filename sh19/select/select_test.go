package _select

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 5000)
	return "Done"
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

func TestSelect(t *testing.T) {

	select {
	case ret := <- AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 1000):
		t.Error("time out")

	}
}








package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 值传递,不用担心会有读取i或重复
		go func(i int) {
			fmt.Println(i)
		}(i)
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 1000)
}

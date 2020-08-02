package fn_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 方法入参和出参都可以是方法
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op * 5
}

func TestFn(t *testing.T) {
	a, b := returnMultiValues()
	t.Log(a, b)

	tsf := timeSpent(slowFun)
	t.Log(tsf(10))
}


func Sum(ops ...int) int{
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3, 4))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func Clear() {
	fmt.Println("Clear resources.")
}

// 延迟执行、必然执行
func TestDefer(t *testing.T) {
	defer Clear()
	fmt.Print("Start")
	panic("err")
}

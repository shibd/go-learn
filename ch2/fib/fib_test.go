package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {

	//var a int = 1
	//var b int = 2

	//var (
	//	a int = 1
	//	b     = 1
	//)
	a := 1
	b := 2
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}

func TestExchange(t *testing.T) {
	a := 1
	b := 2

	// 交换值
	//tmp := a
	//a = b
	//b = tmp

	// 交换值
	a,b = b,a
	t.Log(a, b)

}

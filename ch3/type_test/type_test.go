package type_test

import (
	"testing"
)

type MyInt int64

// go不支持隐性类型转换
func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)

	var c MyInt
	c = MyInt(b)
	t.Log(b, c)
}

// go支持指针,但是有约束
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	// 不支持指针运算
	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)

}

// string的默认初始化是空值，不是nil
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))
	if s == "" {
		t.Log("这是空值")
	}
}
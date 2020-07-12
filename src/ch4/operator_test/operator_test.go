package operator_test

import "testing"

const (
	Readable = 1 << iota // 0001
	Writable // 0010
	Executable // 0100
)

// 数组可以直接比较
func TestCompareArray(t *testing.T) {

	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 2}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}

	t.Log(a == b)
	// 长度不同的数组是编译错误
	//t.Log(a == c)
	t.Log(a == d)
}

// 按位运算清0  &^
func TestBitClear(t *testing.T) {
	a := 7 // 0111
	t.Logf("%b", a)
	t.Logf("%b", Readable)
	t.Logf("%b", Writable)
	t.Logf("%b", Executable)
	a = a &^ Readable
	t.Logf("%b", a)
	a = a &^ Writable
	t.Logf("%b", a)
	a = a &^ Executable
	t.Logf("%b", a)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
}
package string

import "testing"

func TestString(t *testing.T) {

	var s string
	t.Log(s)
	s = "hello"
	t.Log(len(s))

	//s[1] = '3' // string是不可变的

	s = "\xE4\xB8\xA5"
	//s = "\xE4\xBA\xBB\xFF"
	t.Log(s)
	t.Log(len(s)) // byte的长度

	s = "中"
	t.Log(len(s))

	c := []rune(s)
	t.Log(len(c))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRunne(t *testing.T) {
	s := "中华人民共和国"
	for _, c:= range s {
		t.Logf("%[1]c %[1]d", c)
	}
}

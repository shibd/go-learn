package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {

	switch v := p.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unkonw Type")
	}

}

func TestEmptyInterface(t *testing.T) {

	DoSomething(10)
	DoSomething("10")
}
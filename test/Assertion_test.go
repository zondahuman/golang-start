package test

import (
	"fmt"
	"testing"
)


func judge1(a interface{}) string {
	value, ok := a.(string)
	fmt.Println("value:"+ value + " , ok:" ,ok)
	if !ok{
		fmt.Println("It's not ok for type string")
	}

	return value
}

func Test_Assertion1(t *testing.T) {
	//result := judge("b")
	//fmt.Println("result:"+ result)
	result := judge1("abin")
	fmt.Println("result:"+ result)
}
func Test_Assertion2(t *testing.T) {
	result := judge1("abin")
	fmt.Println("result:"+ result)
}
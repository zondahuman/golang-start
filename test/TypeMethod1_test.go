package test

import (
	"testing"
	"fmt"
)


type binFunc1 func(int, int) int

func add(x, y int) int {
	return x + y
}

func (f binFunc1) Error() string {
	return "binFunc1 error"
}

func Test_Two1(t *testing.T) {
	var err error
	err = binFunc1(add)
	fmt.Println(err)
}

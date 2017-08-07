package test

import (
	"fmt"
	"testing"
)

type A func(int, int)

func (f A)Serve() {
	fmt.Println("serve2")
}

func serve(int,int) {
	fmt.Println("serve1")
}

func Test_TypeMethod1(t *testing.T) {
	a := A(serve)
	a(1,2)
	//a.Serve()
}

package test

import (
	"testing"
	"fmt"
)
// http://www.cnblogs.com/howDo/archive/2013/06/04/GoLang-function.html
// http://www.jianshu.com/p/dbd4e6b4900c

type Human struct{
	name string
	age int
	phone string
}

type Student struct {
	Human
	school string
	loan float32
}

type Employee struct {
	Human
	company string
	monry float32
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

func (h Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func (h Human) Sing(lyrics string){
	fmt.Println("La la la la...", lyrics)
}

func (e Employee) SayHi(lyrics string){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

func Test_One(t *testing.T) {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	//paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	//sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	//Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	var i Men

	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()

}

func Test_sendBeauty(t *testing.T) {

}




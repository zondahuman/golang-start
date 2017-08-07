package test

import (
	"testing"
	"fmt"
	"time"
	"math/rand"
)
// http://www.cnblogs.com/howDo/archive/2013/06/04/GoLang-function.html
// http://www.jianshu.com/p/dbd4e6b4900c

type binFunc func(int, int) int

func Test_Two(t *testing.T) {
	// seed your random number generator.
	rand.Seed(time.Now().Unix())

	// create a slice of functions
	fns := []binFunc{
		func(x, y int) int { return x + y },
		func(x, y int) int { return x - y },
		func(x, y int) int { return x * y },
		func(x, y int) int { return x / y },
		func(x, y int) int { return x % y },
	}

	// pick one of those functions at random
	fn := fns[rand.Intn(len(fns))]

	// and execute it
	x, y := 12, 5
	fmt.Println(fn(x, y))
}




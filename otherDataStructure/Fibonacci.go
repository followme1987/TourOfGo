package otherDataStructure

import "fmt"

func fibonacci() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

//Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
func Fib() {
	fun := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ,", fun())
	}
}
